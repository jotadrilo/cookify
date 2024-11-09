package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/core/domain"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID int64 `bun:"id,pk,autoincrement"`

	UUID           string          `bun:"uuid,pk"`
	Name           string          `bun:"name"`
	LangEsES       string          `bun:"lang_es_es"`
	LangEnUS       string          `bun:"lang_en_us"`
	Vendor         string          `bun:"vendor"`
	Unit           string          `bun:"unit"`
	NutritionFacts *NutritionFacts `bun:"rel:has-one,join:uuid=product_uuid"`
}

var _ bun.BeforeAppendModelHook = (*Product)(nil)

func (r *Product) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.UUID == "" {
		r.UUID = uuid.NewString()
	}

	return nil
}

func ProductToDomainProduct(x *Product) *domain.Product {
	if x == nil {
		return nil
	}

	var product = &domain.Product{
		UUID:           x.UUID,
		Name:           x.Name,
		LangEsES:       x.LangEsES,
		LangEnUS:       x.LangEnUS,
		Vendor:         x.Vendor,
		Unit:           domain.ParseUnit(x.Unit),
		NutritionFacts: NutritionFactsToDomainNutritionFacts(x.NutritionFacts),
	}

	if product.NutritionFacts != nil {
		product.NutritionFacts.Product = product
	}

	return product
}

func DomainProductToProduct(x *domain.Product) *Product {
	if x == nil {
		return nil
	}

	var product = &Product{
		UUID:           x.UUID,
		Name:           x.Name,
		LangEsES:       x.LangEsES,
		LangEnUS:       x.LangEnUS,
		Vendor:         x.Vendor,
		Unit:           x.Unit.String(),
		NutritionFacts: DomainNutritionFactsToNutritionFacts(x.NutritionFacts),
	}

	return product
}
