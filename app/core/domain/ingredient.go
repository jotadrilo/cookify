package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Ingredient struct {
	Product  *Product
	Quantity float32
}

func (x *Ingredient) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Product, validation.Required),
		validation.Field(&x.Quantity, validation.Required),
	)
}
