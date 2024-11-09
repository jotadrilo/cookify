package services

import (
	"context"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
)

type ProductsService interface {
	CreateProduct(context.Context, *domain.Product) (*domain.Product, error)
	ListProducts(context.Context) ([]*domain.Product, error)
	GetProductByID(context.Context, string) (*domain.Product, error)
}

type ProductsServiceUnimpl struct{}

func (x *ProductsServiceUnimpl) CreateProduct(context.Context, *domain.Product) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("CreateProduct")
}

func (x *ProductsServiceUnimpl) ListProducts(context.Context) ([]*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("ListProducts")
}

func (x *ProductsServiceUnimpl) GetProductByID(context.Context, string) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("GetProductByID")
}
