package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/adapters/repositories/pg/model"
	"github.com/jotadrilo/cookify/app/adapters/repositories/unimpl"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type ProductsRepository struct {
	unimpl.ProductsRepository

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
	if pp, err := x.GetProductBySlug(ctx, p.Slug); err == nil {
		return pp, nil
	}

	var product = model.DomainProductToProduct(p)

	_, err := x.db.NewInsert().
		Model(product).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	var (
		inserted = model.ProductToDomainProduct(product)
		facts    = model.DomainNutritionFactsToNutritionFacts(inserted.NutritionFacts)
	)

	facts.ProductUUID = product.UUID

	if _, err := x.db.NewInsert().
		Model(facts).
		Exec(ctx); err != nil {
		return nil, err
	}

	inserted.NutritionFacts = model.NutritionFactsToDomainNutritionFacts(facts)

	return inserted, nil
}

func (x *ProductsRepository) ListProducts(ctx context.Context) ([]*domain.Product, error) {
	var (
		items    []*domain.Product
		products []*model.Product
		sel      = x.db.NewSelect().Model(&products).Relation("NutritionFacts")
	)

	if err := sel.Scan(ctx); err != nil {
		return nil, err
	}

	for _, p := range products {
		items = append(items, model.ProductToDomainProduct(p))
	}

	return items, nil
}

func (x *ProductsRepository) GetProductByUUID(ctx context.Context, uuid string) (*domain.Product, error) {
	var (
		product model.Product
		sel     = x.db.NewSelect().Model(&product).Where("uuid = ?", uuid)
	)

	if err := sel.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("product")
		}
		return nil, err
	}

	return model.ProductToDomainProduct(&product), nil
}

func (x *ProductsRepository) GetProductBySlug(ctx context.Context, name string) (*domain.Product, error) {
	var (
		product model.Product
		sel     = x.db.NewSelect().Model(&product).Where("name = ?", name)
	)

	if err := sel.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("product")
		}
		return nil, err
	}

	return model.ProductToDomainProduct(&product), nil
}

func (x *ProductsRepository) GetProductNutritionFactsByUUID(ctx context.Context, uuid string) (*domain.NutritionFacts, error) {
	var (
		product model.Product
		q       = x.db.NewSelect().Model(&product).Where("product.uuid = ?", uuid).Relation("NutritionFacts")
	)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("product")
		}
		return nil, err
	}

	return model.NutritionFactsToDomainNutritionFacts(product.NutritionFacts), nil
}
