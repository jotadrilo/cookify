package fs

import (
	"context"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/jotadrilo/cookify/adapters/repositories/fs/model"
	"github.com/jotadrilo/cookify/adapters/repositories/unimpl"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
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
	if vv, err := x.GetProductByName(ctx, v.Name); err == nil {
		return vv, nil
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

func (x *ProductsRepository) GetProductByName(_ context.Context, productName string) (*domain.Product, error) {
	items, err := decodeJSON[*model.Product](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Name == productName {
			return model.ProductToDomainProduct(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("product")
}

func (x *ProductsRepository) GetProductNutritionFactsByUUID(ctx context.Context, productID string) (*domain.NutritionFacts, error) {
	if vv, err := x.GetProductByUUID(ctx, productID); err == nil {
		return vv.NutritionFacts, nil
	}

	return nil, errorutils.NewErrNotFound("product")
}
