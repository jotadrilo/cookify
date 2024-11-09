package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID int64 `bun:"id,pk,autoincrement"`

	UUID                string          `bun:"uuid,pk"`
	Slug                string          `bun:"slug"`
	DisplayNameLangEsES string          `bun:"display_name_lang_es_es"`
	DisplayNameLangEnUS string          `bun:"display_name_lang_en_us"`
	Vendors             []string        `bun:"vendors"`
	Unit                string          `bun:"unit"`
	NutritionFacts      *NutritionFacts `bun:"rel:has-one,join:uuid=product_uuid"`
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
		UUID:                x.UUID,
		Slug:                x.Slug,
		DisplayNameLangEsES: x.DisplayNameLangEsES,
		DisplayNameLangEnUS: x.DisplayNameLangEnUS,
		Vendors:             x.Vendors,
		Unit:                domain.ParseUnit(x.Unit),
		NutritionFacts:      NutritionFactsToDomainNutritionFacts(x.NutritionFacts),
	}

	return product
}

func DomainProductToProduct(x *domain.Product) *Product {
	if x == nil {
		return nil
	}

	var product = &Product{
		UUID:                x.UUID,
		Slug:                x.Slug,
		DisplayNameLangEsES: x.DisplayNameLangEsES,
		DisplayNameLangEnUS: x.DisplayNameLangEnUS,
		Vendors:             x.Vendors,
		Unit:                x.Unit.String(),
		NutritionFacts:      DomainNutritionFactsToNutritionFacts(x.NutritionFacts),
	}

	return product
}
