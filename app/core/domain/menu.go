package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Menu struct {
	UUID                string
	UserUUID            string
	Name                string
	Label               MenuLabel
	Recipes             []*Recipe
	Ingredients         []*Ingredient
	NutritionFactsTotal *NutritionFacts `json:"nutrition_facts_total,omitempty"`
}

func (x *Menu) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Label, validation.Required),
	)
}

func (x *Menu) Fixup() *Menu {
	var facts = &NutritionFacts{}

	for _, recipe := range x.Recipes {
		facts = facts.Sum(recipe.Fixup().NutritionFactsTotal)
	}

	for _, ingredient := range x.Ingredients {
		facts = facts.Sum(ingredient.Product.NutritionFacts.Multiply(ingredient.Quantity / 100))
	}

	x.NutritionFactsTotal = facts

	return x
}
