package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	UUID           string
	Name           string
	LangEsES       string
	LangEnUS       string
	Vendor         string
	Unit           Unit
	NutritionFacts *NutritionFacts
}

func (a *Product) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.LangEsES, validation.Required),
		validation.Field(&a.LangEsES, validation.Required),
		validation.Field(&a.Vendor, validation.Required),
		validation.Field(&a.Unit, validation.Required),
		validation.Field(&a.NutritionFacts, validation.Required),
	)
}
