package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/app/adapters/controllers/gin"
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/client/cmd/config"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/mercadona"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/prompt"
)

var priceRe = regexp.MustCompile(`([\d\\.]+)?\s*(\w+)`)

func newImportMercadonaCmd() *cobra.Command {
	type cmdConfig struct {
		Mercadona config.MercadonaServer
		Cookify   config.CookifyServer
	}

	cfg := &cmdConfig{}

	cmd := &cobra.Command{
		Use:   "mercadona",
		Short: "Import data from mercadona",
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = cmd.Context()

			cli := mercadona.NewClient(cfg.Mercadona.Address, cfg.Mercadona.ZipCode)

			apiCli, err := api.NewClientWithResponses(cfg.Cookify.Address)
			if err != nil {
				return err
			}

			var products []*mercadona.Product

			if cfg.Mercadona.ProductsFile != "" {
				_, err := os.Stat(cfg.Mercadona.ProductsFile)
				if err != nil {
					if !errors.Is(err, os.ErrNotExist) {
						return err
					}
				} else {
					f, err := os.OpenFile(cfg.Mercadona.ProductsFile, os.O_RDONLY, os.ModePerm)
					if err != nil {
						return err
					}

					if err := json.NewDecoder(f).Decode(&products); err != nil {
						return err
					}
				}
			}

			if len(products) == 0 {
				logger.Infof("Loading products from Mercadona API... This may take a while")

				products, err = cli.GetProducts(ctx, &mercadona.GetProductsOptions{IgnoreCategoriesWithPrefix: cfg.Mercadona.IgnoreCategoriesWithPrefix})
				if err != nil {
					return err
				}

				logger.Infof("Found %d products", len(products))

				if cfg.Mercadona.ProductsFile != "" {
					f, err := os.Create(cfg.Mercadona.ProductsFile)
					if err != nil {
						return err
					}

					enc := json.NewEncoder(f)
					enc.SetIndent("", "  ")

					if err := enc.Encode(products); err != nil {
						return err
					}
				}
			}

			var (
				slugs         []string
				productBySlug = make(map[string]*mercadona.Product)
			)

			for _, prod := range products {
				slug := fmt.Sprintf("%s (%s - %s)", prod.DisplayName, prod.ID, prod.Slug)
				slugs = append(slugs, slug)
				productBySlug[slug] = prod
			}

			sort.Strings(slugs)

			var selectedSlugs = cfg.Mercadona.TargetProductSlugs

			if len(selectedSlugs) == 0 {
				prompt.MustAsk(&selectedSlugs, prompt.WithAskPrompt(&survey.MultiSelect{
					Message:  "Select the target products:",
					Options:  slugs,
					PageSize: 50,
				}), prompt.WithAskKeepFilter())
			}

			var selectedProducts []*mercadona.Product

			for _, slug := range selectedSlugs {
				logger.Infof("Fetching %q product", slug)

				product, err := cli.GetProduct(ctx, productBySlug[slug].ID)
				if err != nil {
					return err
				}

				selectedProducts = append(selectedProducts, product)

				logger.Infof("Product: %s", product.StringJSON())

				price, err := strconv.ParseFloat(product.Price.ReferencePrice, 32)
				if err != nil {
					return err
				}

				matches := priceRe.FindStringSubmatch(product.Price.ReferenceFormat)

				var uh domain.UnitHelper
				if matches[1] == "" {
					uh = domain.ParseUnitHelper(matches[2], 1)
				} else {
					q, err := strconv.ParseFloat(matches[1], 32)
					if err != nil {
						return err
					}
					uh = domain.ParseUnitHelper(matches[2], float32(q))
				}

				var body = api.PostProductsJSONRequestBody{
					DisplayNameLangEsEs: oapi.String(product.DisplayName),
					Slug:                product.Slug,
					Unit:                gin.DomainUnitToAPIUnit(uh.Unit()),
					CurrentPrice: &api.Price{
						Price:    float32(price),
						Quantity: uh.Value(),
					},
					Vendors: oapi.StringSlice([]string{product.Brand}),
				}

				bindNutritionValues(&body)

				rsp, err := apiCli.PostProductsWithResponse(ctx, body)
				if err != nil {
					return err
				}

				switch rsp.StatusCode() {
				case http.StatusCreated:
					var id api.ProductID
					if err := json.Unmarshal(rsp.Body, &id); err != nil {
						return err
					}
					logger.Infof("Product %q created: %s", body.Slug, id.Uuid.String())
				case http.StatusConflict:
					var id api.ProductID
					if err := json.Unmarshal(rsp.Body, &id); err == nil {
						logger.Infof("Product %q already exists: %s", body.Slug, id.Uuid.String())
						break
					}

					var apiError api.Error
					if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
						return err
					}
					logger.Warnf("Product %q was not created: (%d - %s - %s)", body.Slug, rsp.StatusCode(), rsp.Status(), apiError.Title)
				default:
					var apiError api.Error
					if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
						return err
					}
					return fmt.Errorf("product %q was not created: (%d - %s - %s)", body.Slug, rsp.StatusCode(), rsp.Status(), apiError.Title)
				}

			}

			return nil
		},
	}

	config.BindMercadonaServerFlags(cmd, &cfg.Mercadona)
	config.BindCookifyServerFlags(cmd, &cfg.Cookify)

	return cmd
}

func bindNutritionValues(product *api.PostProductsJSONRequestBody) {
	var (
		caffeine           float32
		cal                float32
		calcium            float32
		carbohydrateSugar  float32
		carbohydrateTotal  float32
		cholesterol        float32
		fatMonounsaturated float32
		fatPolyunsaturated float32
		fatSaturated       float32
		fatTotal           float32
		fiber              float32
		iron               float32
		potassium          float32
		protein            float32
		salt               float32
		sodium             float32
		vitaminA           float32
		vitaminB1          float32
		vitaminB12         float32
		vitaminB2          float32
		vitaminB3          float32
		vitaminB4          float32
		vitaminB5          float32
		vitaminB6          float32
		vitaminC           float32
		vitaminD           float32
		vitaminE           float32
		vitaminK           float32
	)

	var ask = func(v any, what string) {
		prompt.MustAskFloat32(v, prompt.WithAskPrompt(&survey.Input{
			Message: fmt.Sprintf("Enter the product %s:", what),
			Default: "0.0",
		}))
	}

	ask(&cal, "calories")
	ask(&fatTotal, "total fats")
	ask(&fatSaturated, "sat fats")
	ask(&carbohydrateTotal, "total carbs")
	ask(&carbohydrateSugar, "sugar carbs")
	ask(&fiber, "fiber")
	ask(&protein, "protein")
	ask(&salt, "salt")

	if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Enter others?")) {
		ask(&fatMonounsaturated, "mono fats")
		ask(&fatPolyunsaturated, "poly fats")
		ask(&cholesterol, "cholesterol")
		ask(&calcium, "calcium")
		ask(&iron, "iron")
		ask(&caffeine, "caffeine")
		ask(&potassium, "potassium")
		ask(&sodium, "sodium")
	}

	if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Enter vitamins?")) {
		ask(&vitaminA, "vitnamin A")
		ask(&vitaminB1, "vitnamin B1")
		ask(&vitaminB2, "vitnamin B2")
		ask(&vitaminB3, "vitnamin B3")
		ask(&vitaminB4, "vitnamin B4")
		ask(&vitaminB5, "vitnamin B5")
		ask(&vitaminB6, "vitnamin B6")
		ask(&vitaminB12, "vitnamin B12")
		ask(&vitaminC, "vitnamin C")
		ask(&vitaminD, "vitnamin D")
		ask(&vitaminE, "vitnamin E")
		ask(&vitaminK, "vitnamin K")
	}

	var toFloatPtr = func(f float32) *float32 {
		if f == 0 {
			return nil
		}
		var ff = float32(f)
		return &ff
	}

	product.NutritionFacts100 = api.NutritionFacts{
		Caffeine:           toFloatPtr(caffeine),
		Cal:                float32(cal * 1000),
		Calcium:            toFloatPtr(calcium),
		CarbohydrateSugar:  toFloatPtr(carbohydrateSugar),
		CarbohydrateTotal:  toFloatPtr(carbohydrateTotal),
		Cholesterol:        toFloatPtr(cholesterol),
		FatMonounsaturated: toFloatPtr(fatMonounsaturated),
		FatPolyunsaturated: toFloatPtr(fatPolyunsaturated),
		FatSaturated:       toFloatPtr(fatSaturated),
		FatTotal:           toFloatPtr(fatTotal),
		Fiber:              toFloatPtr(fiber),
		Iron:               toFloatPtr(iron),
		Potassium:          toFloatPtr(potassium),
		Protein:            toFloatPtr(protein),
		Salt:               toFloatPtr(salt),
		Sodium:             toFloatPtr(sodium),
		VitaminA:           toFloatPtr(vitaminA),
		VitaminB1:          toFloatPtr(vitaminB1),
		VitaminB12:         toFloatPtr(vitaminB12),
		VitaminB2:          toFloatPtr(vitaminB2),
		VitaminB3:          toFloatPtr(vitaminB3),
		VitaminB4:          toFloatPtr(vitaminB4),
		VitaminB5:          toFloatPtr(vitaminB5),
		VitaminB6:          toFloatPtr(vitaminB6),
		VitaminC:           toFloatPtr(vitaminC),
		VitaminD:           toFloatPtr(vitaminD),
		VitaminE:           toFloatPtr(vitaminE),
		VitaminK:           toFloatPtr(vitaminK),
	}

	if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Is it incorrect?")) {
		bindNutritionValues(product)
	}
}
