package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Recipe struct {
	UUID                string
	UserUUID            string
	Name                string
	Ingredients         []*Ingredient
	Quantity            float32
	NutritionFacts      *NutritionFacts
	NutritionFactsTotal *NutritionFacts
}

func (x *Recipe) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Ingredients, validation.Required),
	)
}

func (x *Recipe) Fixup() *Recipe {
	var (
		facts            = &NutritionFacts{}
		quantity float32 = 0
	)

	// https://diagon.arthursonzogni.com/#Math
	//
	// sum(Recipe_totalFacts,n=0,N) = (Product(n)_facts * Product(n)_quantity / 100)
	//
	//     N                      ⎛Product(n)      ⋅ Product(n)        ⎞
	//    ___                     ⎜          facts             quantity⎟
	//    ╲    Recipe           = ⎜────────────────────────────────────⎟
	//    ╱          totalFacts   ⎝                 100                ⎠
	//    ‾‾‾
	//    n = 0
	//
	// sum(Recipe_quantity,n=0,N) = Product(n)_quantity
	//
	//     N
	//    ___
	//    ╲    Recipe         = Product(n)
	//    ╱          quantity             quantity
	//    ‾‾‾
	//    n = 0
	//
	// Recipe_facts = Recipe_totalFacts * 100 / Recipe_quantity
	//
	//                  Recipe           ⋅ 100
	//                        totalFacts
	//    Recipe      = ──────────────────────
	//          facts       Recipe
	//                            quantity
	//

	for _, ingredient := range x.Ingredients {
		productFacts := ingredient.Product.NutritionFacts.Multiply(ingredient.Quantity / 100)
		facts = facts.Sum(productFacts)
		quantity += ingredient.Quantity
	}

	x.Quantity = quantity
	x.NutritionFactsTotal = facts
	x.NutritionFacts = facts.Multiply(100 / quantity)

	return x
}
