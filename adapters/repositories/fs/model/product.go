package model

import "github.com/jotadrilo/cookify/core/domain"

type Product struct {
	UUID           string          `json:"uuid"`
	Name           string          `json:"name"`
	LangEsES       string          `json:"lang_es_es"`
	LangEnUS       string          `json:"lang_en_us"`
	Vendor         string          `json:"vendor"`
	Unit           string          `json:"unit"`
	NutritionFacts *NutritionFacts `json:"nutrition_facts"`
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
