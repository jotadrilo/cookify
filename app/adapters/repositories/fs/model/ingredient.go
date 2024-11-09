package model

import "github.com/jotadrilo/cookify/app/core/domain"

type Ingredient struct {
	Product  *Product `json:"product"`
	Quantity float32  `json:"quantity"`
}

func IngredientToDomainIngredient(x *Ingredient) *domain.Ingredient {
	if x == nil {
		return nil
	}

	var ingredient = &domain.Ingredient{
		Product:  ProductToDomainProduct(x.Product),
		Quantity: x.Quantity,
	}

	return ingredient
}

func DomainIngredientToIngredient(x *domain.Ingredient) *Ingredient {
	if x == nil {
		return nil
	}

	return &Ingredient{
		Product:  DomainProductToProduct(x.Product),
		Quantity: x.Quantity,
	}
}
