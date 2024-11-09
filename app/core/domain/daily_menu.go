package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DailyMenu struct {
	UUID                string
	UserUUID            string
	Name                string
	Menus               []*Menu
	NutritionFactsTotal *NutritionFacts
}

func (x *DailyMenu) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Name, validation.Required),
	)
}

func (x *DailyMenu) Fixup() *DailyMenu {
	var facts = &NutritionFacts{}

	for _, menu := range x.Menus {
		facts = facts.Sum(menu.Fixup().NutritionFactsTotal)
	}

	x.NutritionFactsTotal = facts

	return x
}
