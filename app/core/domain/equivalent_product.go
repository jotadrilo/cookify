package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type EquivalentProduct struct {
	Score              float32
	Product            *Product
	NutritionFactsDiff *NutritionFacts
}

func (a *EquivalentProduct) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Score, validation.Required),
		validation.Field(&a.Product, validation.Required),
		validation.Field(&a.NutritionFactsDiff, validation.Required),
	)
}
