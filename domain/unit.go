package domain

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Unit struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Unit10em3 string `json:"unit_10em3"`
	Unit1     string `json:"unit_1"`
	Unit10e3  string `json:"unit_10e3"`
}

func (a Unit) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Unit1, validation.Required),
	)
}

func (x *Unit) String() string {
	if x == nil {
		return ""
	}

	return x.Name
}

func (x *Unit) GetQuantityString(q float32) string {
	if x == nil {
		return ""
	}

	if q < 10e-3 && x.Unit10em3 != "" {
		return fmt.Sprintf("%.2f%s", q*10e3, x.Unit10em3)
	}

	if q > 10e3 && x.Unit10e3 != "" {
		return fmt.Sprintf("%.2f%s", q*10e-3, x.Unit10e3)
	}

	return fmt.Sprintf("%.2f%s", q, x.Unit1)
}
