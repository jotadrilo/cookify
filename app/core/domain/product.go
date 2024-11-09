package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	UUID                string
	Slug                string
	DisplayNameLangEsES string
	DisplayNameLangEnUS string
	Vendors             []string
	CurrentPrice        *Price
	Unit                Unit
	NutritionFacts      *NutritionFacts
}

func (a *Product) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Slug, validation.Required),
		validation.Field(&a.DisplayNameLangEsES, validation.Required),
		validation.Field(&a.Unit, validation.Required),
	)
}
