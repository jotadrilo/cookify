package services

import (
	"context"
	"fmt"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/ports/repositories"
	"github.com/jotadrilo/cookify/ports/services"
)

type ProductsService struct {
	services.ProductsServiceUnimpl

	Products repositories.ProductsRepository
}

type ProductsServiceOptions struct {
	Products repositories.ProductsRepository
}

func NewProductsService(opts *ProductsServiceOptions) *ProductsService {
	return &ProductsService{
		Products: opts.Products,
	}
}

func (x *ProductsService) CreateProduct(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	pn, err := x.Products.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created product %q", pn.ID)

	return pn, nil
}

func (x *ProductsService) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	return x.Products.ListProducts(ctx)
}

func (x *ProductsService) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	return x.Products.GetProductByID(ctx, id)
}
