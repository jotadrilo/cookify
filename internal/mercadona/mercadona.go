package mercadona

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/slices"
)

type Client struct {
	Address   string
	Client    *http.Client
	ZipCode   uint32
	Warehouse string
	once      sync.Once
}

func NewClient(address string, zip uint32) *Client {
	return &Client{
		Address: address,
		Client:  &http.Client{},
		ZipCode: zip,
	}
}

type Categories struct {
	Results []*Category
}

type Category struct {
	ID         uint32      `json:"id"`
	Name       string      `json:"name"`
	Categories []*Category `json:"categories,omitempty"`
	Products   []*Product  `json:"products,omitempty"`
}

func (x *Category) String() string {
	var parts []string

	parts = append(parts, fmt.Sprintf("id: %d", x.ID))
	parts = append(parts, fmt.Sprintf("name: %s", x.Name))
	parts = append(parts, fmt.Sprintf("categories: %d", len(x.Categories)))
	parts = append(parts, fmt.Sprintf("products: %d", len(x.Products)))

	return fmt.Sprintf("Category(%s)", strings.Join(parts, ", "))
}

type Product struct {
	ID          string   `json:"id"`
	Slug        string   `json:"slug"`
	DisplayName string   `json:"display_name"`
	Brand       string   `json:"brand"`
	Price       *Price   `json:"price_instructions,omitempty"`
	Photos      []*Photo `json:"photos,omitempty"`
}

func (x *Product) String() string {
	var parts []string

	if x.ID != "" {
		parts = append(parts, fmt.Sprintf("id: %s", x.ID))
	}

	if x.DisplayName != "" {
		parts = append(parts, fmt.Sprintf("display_name: %s", x.DisplayName))
	}

	if x.Price != nil {
		parts = append(parts, fmt.Sprintf("reference_price: %s", x.Price.ReferencePrice))
		parts = append(parts, fmt.Sprintf("reference_format: %s", x.Price.ReferenceFormat))
	}

	return fmt.Sprintf("Product(%s)", strings.Join(parts, ", "))
}

func (x *Product) StringJSON() string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetIndent("", "  ")
	_ = enc.Encode(x)
	return b.String()
}

type Price struct {
	ReferencePrice  string `json:"reference_price"`
	ReferenceFormat string `json:"reference_format"`
}

type Photo struct {
	Zoom string `json:"zoom"`
}

func (x *Client) lookupWarehouse() error {
	if x.Warehouse != "" || x.ZipCode == 0 {
		return nil
	}

	queryURL, err := url.Parse(x.Address)
	if err != nil {
		return err
	}

	queryURL.Path += "/postal-codes/actions/change-pc/"

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(struct {
		ZipCode string `json:"new_postal_code"`
	}{
		ZipCode: fmt.Sprintf("%d", x.ZipCode),
	}); err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), &b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	rsp, err := x.Client.Do(req)
	if err != nil {
		return err
	}

	x.Warehouse = rsp.Header.Get("x-customer-wh")

	return nil
}

func (x *Client) doGET(ctx context.Context, path string, v any) error {
	x.once.Do(func() {
		if err := x.lookupWarehouse(); err != nil {
			logger.Warnf("Cannot lookup near warehouse. Results may not be accurate: %s", err.Error())
		}
	})

	queryURL, err := url.Parse(x.Address)
	if err != nil {
		return err
	}

	queryURL.Path += path

	if x.Warehouse != "" {
		queryURL.RawQuery = fmt.Sprintf("wh=%s", x.Warehouse)
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return err
	}

	req.WithContext(ctx)

	req.Header.Add("Accept", "application/json")

	logger.Debugf("Doing GET to %s", req.URL.String())
	rsp, err := x.Client.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = rsp.Body.Close() }()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	switch {
	case rsp.StatusCode == 200:
		if err := json.Unmarshal(bodyBytes, &v); err != nil {
			return err
		}
	default:
		return err
	}

	return nil
}

func (x *Client) ListCategories(ctx context.Context) ([]*Category, error) {
	var resp Categories

	logger.Debugf("Fetching categories")

	if err := x.doGET(ctx, "/categories/", &resp); err != nil {
		return nil, err
	}

	return flattenCategories(resp.Results, nil), nil
}

func (x *Client) GetCategory(ctx context.Context, id uint32) (*Category, error) {
	var resp = &Category{}

	if err := x.doGET(ctx, fmt.Sprintf("/categories/%d/", id), resp); err != nil {
		return nil, err
	}

	logger.Debugf("Fetched category %s (%d)", resp.Name, id)

	return resp, nil
}

func (x *Client) ListCategoryProducts(ctx context.Context, id uint32) ([]*Product, error) {
	category, err := x.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	var products []*Product

	for _, cat := range category.Categories {
		for _, prod := range cat.Products {
			products = append(products, &Product{
				ID:          prod.ID,
				Slug:        prod.Slug,
				DisplayName: prod.DisplayName,
			})
		}
	}

	for _, prod := range category.Products {
		products = append(products, &Product{
			ID:          prod.ID,
			Slug:        prod.Slug,
			DisplayName: prod.DisplayName,
		})
	}

	return products, nil
}

type GetProductsOptions struct {
	IgnoreCategoriesWithPrefix []string
}

func (x *Client) GetProducts(ctx context.Context, opts *GetProductsOptions) ([]*Product, error) {
	var products []*Product

	categories, err := x.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		if opts != nil {
			if slices.Any(opts.IgnoreCategoriesWithPrefix, func(s string) bool {
				return strings.HasPrefix(cat.Name, s)
			}) {
				continue
			}
		}

		catProducts, err := x.ListCategoryProducts(ctx, cat.ID)
		if err != nil {
			return nil, err
		}
		products = append(products, catProducts...)
	}

	return products, nil
}

func (x *Client) GetProduct(ctx context.Context, id string) (*Product, error) {
	var resp = &Product{}

	if err := x.doGET(ctx, fmt.Sprintf("/products/%s/", id), resp); err != nil {
		return nil, err
	}

	logger.Debugf("Fetched product %s (%s)", resp.Slug, id)

	return resp, nil
}

func flattenCategories(s []*Category, parent *Category) []*Category {
	var categories []*Category

	for _, cat := range s {
		if parent != nil {
			cat.Name = fmt.Sprintf("%s > %s", parent.Name, cat.Name)
		}

		if len(cat.Categories) == 0 {
			categories = append(categories, cat)
			continue
		}

		categories = append(categories, flattenCategories(cat.Categories, cat)...)
	}

	return categories
}
