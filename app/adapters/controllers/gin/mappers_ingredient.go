package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/slices"
)

func DomainIngredientToAPIIngredient(v *domain.Ingredient) api.Ingredient {
	if v == nil {
		return api.Ingredient{}
	}
	return api.Ingredient{
		Product:  *DomainProductToAPIProduct(v.Product),
		Quantity: v.Quantity,
	}
}

func DomainIngredientsToAPIIngredients(v []*domain.Ingredient) []api.Ingredient {
	return slices.Map(v, DomainIngredientToAPIIngredient)
}

func IngredientToDomainIngredient(v *api.Ingredient) *domain.Ingredient {
	if v == nil {
		return &domain.Ingredient{}
	}
	return &domain.Ingredient{
		Quantity: v.Quantity,
		Product:  ProductToDomainProduct(&v.Product),
	}
}

func IngredientsToDomainIngredients(s []api.Ingredient) []*domain.Ingredient {
	if s == nil {
		return nil
	}

	var ingredients []*domain.Ingredient

	for _, v := range s {
		ingredients = append(ingredients, IngredientToDomainIngredient(&v))
	}

	return ingredients
}
