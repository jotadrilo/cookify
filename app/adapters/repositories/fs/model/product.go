package model

import (
	"github.com/jotadrilo/cookify/app/core/domain"
)

type Product struct {
	UUID                string          `json:"uuid"`
	Slug                string          `json:"slug"`
	DisplayNameLangEsES string          `json:"display_name_lang_es_es"`
	DisplayNameLangEnUS string          `json:"display_name_lang_en_us,omitempty"`
	Vendors             []string        `json:"vendors,omitempty"`
	Unit                string          `json:"unit"`
	CurrentPrice        *Price          `json:"price,omitempty"`
	NutritionFacts      *NutritionFacts `json:"nutrition_facts,omitempty"`
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
		CurrentPrice:        PriceToDomainPrice(x.CurrentPrice),
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
		CurrentPrice:        DomainPriceToPrice(x.CurrentPrice),
		NutritionFacts:      DomainNutritionFactsToNutritionFacts(x.NutritionFacts),
	}

	return product
}
