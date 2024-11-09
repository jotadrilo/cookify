package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Units []*Unit `json:"units,omitempty"`
}

func (a Product) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Units, validation.Required),
	)
}
