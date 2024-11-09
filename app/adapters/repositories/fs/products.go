package fs

import (
	"context"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/jotadrilo/cookify/app/adapters/repositories/fs/model"
	"github.com/jotadrilo/cookify/app/adapters/repositories/unimpl"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/slices"
)

type ProductsRepository struct {
	unimpl.ProductsRepository

	root string
}

var _ ports.ProductsRepository = (*ProductsRepository)(nil)

type ProductsRepositoryOptions struct {
	Root string
}

func NewProductsRepository(opts *ProductsRepositoryOptions) *ProductsRepository {
	return &ProductsRepository{
		root: opts.Root,
	}
}

func (x *ProductsRepository) getFile() string {
	return filepath.Join(x.root, "products.json")
}

func (x *ProductsRepository) CreateProduct(ctx context.Context, v *domain.Product) (*domain.Product, error) {
	if vv, err := x.GetProductBySlug(ctx, v.Slug); err == nil {
		return vv, errorutils.NewErrAlreadyExists("product")
	}

	var (
		vv = model.DomainProductToProduct(v)
	)

	vv.UUID = uuid.New().String()

	if err := appendToJSON(x.getFile(), vv); err != nil {
		return nil, err
	}

	return model.ProductToDomainProduct(vv), nil
}

func (x *ProductsRepository) ListProducts(_ context.Context) ([]*domain.Product, error) {
	items, err := decodeJSON[*model.Product](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map[*model.Product, *domain.Product](items, model.ProductToDomainProduct), nil
}

func (x *ProductsRepository) GetProductByUUID(_ context.Context, productID string) (*domain.Product, error) {
	items, err := decodeJSON[*model.Product](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == productID {
			return model.ProductToDomainProduct(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("product")
}

func (x *ProductsRepository) GetProductBySlug(_ context.Context, productSlug string) (*domain.Product, error) {
	items, err := decodeJSON[*model.Product](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Slug == productSlug {
			return model.ProductToDomainProduct(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("product")
}
