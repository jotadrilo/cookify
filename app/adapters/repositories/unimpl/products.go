package unimpl

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type ProductsRepository struct {
	ports.Repository
}

var _ ports.ProductsRepository = (*ProductsRepository)(nil)

func (x *ProductsRepository) CreateProduct(context.Context, *domain.Product) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("CreateProduct")
}

func (x *ProductsRepository) ListProducts(context.Context) ([]*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("ListProducts")
}

func (x *ProductsRepository) GetProductByUUID(context.Context, string) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("GetProductByUUID")
}

func (x *ProductsRepository) GetProductNutritionFactsByUUID(context.Context, string) (*domain.NutritionFacts, error) {
	return nil, errorutils.NewErrNotImplemented("GetProductNutritionFactsByUUID")
}
