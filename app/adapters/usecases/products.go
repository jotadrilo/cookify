package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
)

type UnimplementedProductsUseCase struct {
	ports.UseCase
}

var _ ports.ProductsUseCase = (*UnimplementedProductsUseCase)(nil)

func (x *UnimplementedProductsUseCase) CreateProduct(context.Context, *domain.Product) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("CreateProduct")
}

func (x *UnimplementedProductsUseCase) ListProducts(context.Context) ([]*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("ListProducts")
}

func (x *UnimplementedProductsUseCase) GetProductByUUID(context.Context, string) (*domain.Product, error) {
	return nil, errorutils.NewErrNotImplemented("GetProductByUUID")
}

type ProductsUseCase struct {
	UnimplementedProductsUseCase

	Products ports.ProductsRepository
}

type ProductsUseCaseOptions struct {
	Products ports.ProductsRepository
}

func NewProductsUseCase(opts *ProductsUseCaseOptions) *ProductsUseCase {
	return &ProductsUseCase{
		Products: opts.Products,
	}
}

func (x *ProductsUseCase) CreateProduct(ctx context.Context, v *domain.Product) (*domain.Product, error) {
	vv, err := x.Products.CreateProduct(ctx, v)
	if err != nil {
		logger.Errorf("Cannot create product %v: %s", v, err.Error())

		if errors.Is(err, errorutils.ErrAlreadyExists) {
			return nil, errorutils.NewErrAlreadyExists(fmt.Sprintf("product %s", vv.UUID))
		}

		return nil, errorutils.NewErrNotCreated("product")
	}

	logger.Infof("Created product %q", vv.UUID)

	return vv, nil
}

func (x *ProductsUseCase) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	return x.Products.ListProducts(ctx)
}

func (x *ProductsUseCase) GetProductByUUID(ctx context.Context, uuid string) (*domain.Product, error) {
	return x.Products.GetProductByUUID(ctx, uuid)
}
