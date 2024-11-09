package repositories

import (
	"context"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
)

type ProductsRepository interface {
	CreateProduct(context.Context, *domain.Product) (*domain.Product, error)
	ListProducts(context.Context) ([]*domain.Product, error)
	GetProductByID(context.Context, string) (*domain.Product, error)
}

type ProductsRepositoryUnimpl struct{}

func (x *ProductsRepositoryUnimpl) CreateProduct(context.Context, *domain.Product) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("CreateProduct")
}

func (x *ProductsRepositoryUnimpl) ListProducts(context.Context) ([]*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("ListProducts")
}

func (x *ProductsRepositoryUnimpl) GetProductByID(context.Context, string) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("GetProductByID")
}
