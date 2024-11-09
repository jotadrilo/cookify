package pg

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jotadrilo/cookify/adapters/repositories/pg/model"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
	"github.com/jotadrilo/cookify/ports/repositories"
	"github.com/uptrace/bun"
)

type ProductsRepository struct {
	repositories.ProductsRepositoryUnimpl

	db *bun.DB
}

type ProductsRepositoryOptions struct {
	DB *bun.DB
}

func NewProductsRepository(opts *ProductsRepositoryOptions) *ProductsRepository {
	return &ProductsRepository{
		db: opts.DB,
	}
}

func (x *ProductsRepository) CreateProduct(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	p.ID = uuid.NewString()

	var (
		q = x.db.NewInsert().Model(model.DomainProductToProduct(p))
	)

	if _, err := q.Exec(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

func (x *ProductsRepository) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	var (
		items    []*domain.Product
		products []*model.Product
		q        = x.db.NewSelect().Model(&products)
	)

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	for _, p := range products {
		items = append(items, model.ProductToDomainProduct(p))
	}

	return items, nil
}

func (x *ProductsRepository) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	var (
		product model.Product
		q       = x.db.NewSelect().Model(&product).Where("uuid = ?", id)
	)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("product")
		}
		return nil, err
	}

	return model.ProductToDomainProduct(&product), nil
}
