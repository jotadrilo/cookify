package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Price struct {
	UUID     string
	Price    float32
	Quantity float32
}

func (a *Price) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Price, validation.Required),
		validation.Field(&a.Quantity, validation.Required),
	)
}
