package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/slices"
)

func DomainProductToAPIProduct(v *domain.Product) *api.Product {
	if v == nil {
		return nil
	}

	return &api.Product{
		Uuid:                oapi.UUID(v.UUID),
		Slug:                v.Slug,
		DisplayNameLangEsEs: oapi.String(v.DisplayNameLangEsES),
		DisplayNameLangEnUs: oapi.String(v.DisplayNameLangEnUS),
		Unit:                DomainUnitToAPIUnit(v.Unit),
		Vendors:             oapi.StringSlice(v.Vendors),
		CurrentPrice:        DomainPriceToAPIPrice(v.CurrentPrice),
		NutritionFacts100:   *DomainNutritionFactsToAPINutritionFacts(v.NutritionFacts),
	}
}

func DomainProductsToAPIProducts(s []*domain.Product) []*api.Product {
	return slices.Map(s, DomainProductToAPIProduct)
}

func ProductToDomainProduct(v *api.Product) *domain.Product {
	if v == nil {
		return nil
	}

	return &domain.Product{
		UUID:                oapi.UUIDValue(v.Uuid),
		Slug:                v.Slug,
		DisplayNameLangEsES: oapi.StringValue(v.DisplayNameLangEsEs),
		DisplayNameLangEnUS: oapi.StringValue(v.DisplayNameLangEnUs),
		Unit:                UnitToDomainUnit(v.Unit),
		Vendors:             oapi.StringSliceValue(v.Vendors),
		CurrentPrice:        PriceToDomainPrice(v.CurrentPrice),
		NutritionFacts:      NutritionFactsToDomainNutritionFacts(&v.NutritionFacts100),
	}
}

func ProductRequestToDomainProduct(v *api.PostProductsJSONRequestBody) *domain.Product {
	if v == nil {
		return nil
	}

	return &domain.Product{
		Slug:                v.Slug,
		DisplayNameLangEsES: oapi.StringValue(v.DisplayNameLangEsEs),
		DisplayNameLangEnUS: oapi.StringValue(v.DisplayNameLangEnUs),
		Unit:                UnitToDomainUnit(v.Unit),
		Vendors:             oapi.StringSliceValue(v.Vendors),
		CurrentPrice:        PriceToDomainPrice(v.CurrentPrice),
		NutritionFacts:      NutritionFactsToDomainNutritionFacts(&v.NutritionFacts100),
	}
}

func DomainEquivalentProductToAPIEquivalentProduct(v *domain.EquivalentProduct) *api.EquivalentProduct {
	if v == nil {
		return nil
	}

	return &api.EquivalentProduct{
		NutritionFacts100Diff: *DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsDiff),
		Product:               *DomainProductToAPIProduct(v.Product),
		Score:                 *oapi.Float32(v.Score),
	}
}

func DomainEquivalentProductsToAPIEquivalentProducts(s []*domain.EquivalentProduct) []*api.EquivalentProduct {
	return slices.Map(s, DomainEquivalentProductToAPIEquivalentProduct)
}
