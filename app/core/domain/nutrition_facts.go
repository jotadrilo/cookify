package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type NutritionFacts struct {
	UUID               string
	Cal                float32
	FatTotal           float32
	FatSaturated       float32
	FatMonounsaturated float32
	FatPolyunsaturated float32
	Cholesterol        float32
	Salt               float32
	Sodium             float32
	Potassium          float32
	CarbohydrateTotal  float32
	CarbohydrateSugar  float32
	Protein            float32
	Fiber              float32
	Calcium            float32
	Iron               float32
	Caffeine           float32
	VitaminA           float32
	VitaminB1          float32
	VitaminB2          float32
	VitaminB3          float32
	VitaminB4          float32
	VitaminB5          float32
	VitaminB6          float32
	VitaminB12         float32
	VitaminC           float32
	VitaminD           float32
	VitaminE           float32
	VitaminK           float32
}

func (a *NutritionFacts) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Cal, validation.Required),
	)
}

func (x *NutritionFacts) Multiply(q float32) *NutritionFacts {
	return &NutritionFacts{
		Cal:                x.Cal * q,
		FatTotal:           x.FatTotal * q,
		FatSaturated:       x.FatSaturated * q,
		FatMonounsaturated: x.FatMonounsaturated * q,
		FatPolyunsaturated: x.FatPolyunsaturated * q,
		Cholesterol:        x.Cholesterol * q,
		Salt:               x.Salt * q,
		Sodium:             x.Sodium * q,
		Potassium:          x.Potassium * q,
		CarbohydrateTotal:  x.CarbohydrateTotal * q,
		CarbohydrateSugar:  x.CarbohydrateSugar * q,
		Protein:            x.Protein * q,
		Fiber:              x.Fiber * q,
		Calcium:            x.Calcium * q,
		Iron:               x.Iron * q,
		Caffeine:           x.Caffeine * q,
		VitaminA:           x.VitaminA * q,
		VitaminB1:          x.VitaminB1 * q,
		VitaminB2:          x.VitaminB2 * q,
		VitaminB3:          x.VitaminB3 * q,
		VitaminB4:          x.VitaminB4 * q,
		VitaminB5:          x.VitaminB5 * q,
		VitaminB6:          x.VitaminB6 * q,
		VitaminB12:         x.VitaminB12 * q,
		VitaminC:           x.VitaminC * q,
		VitaminD:           x.VitaminD * q,
		VitaminE:           x.VitaminE * q,
		VitaminK:           x.VitaminK * q,
	}
}

func (x *NutritionFacts) Sum(y *NutritionFacts) *NutritionFacts {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	return &NutritionFacts{
		Cal:                x.Cal + y.Cal,
		FatTotal:           x.FatTotal + y.FatTotal,
		FatSaturated:       x.FatSaturated + y.FatSaturated,
		FatMonounsaturated: x.FatMonounsaturated + y.FatMonounsaturated,
		FatPolyunsaturated: x.FatPolyunsaturated + y.FatPolyunsaturated,
		Cholesterol:        x.Cholesterol + y.Cholesterol,
		Salt:               x.Salt + y.Salt,
		Sodium:             x.Sodium + y.Sodium,
		Potassium:          x.Potassium + y.Potassium,
		CarbohydrateTotal:  x.CarbohydrateTotal + y.CarbohydrateTotal,
		CarbohydrateSugar:  x.CarbohydrateSugar + y.CarbohydrateSugar,
		Protein:            x.Protein + y.Protein,
		Fiber:              x.Fiber + y.Fiber,
		Calcium:            x.Calcium + y.Calcium,
		Iron:               x.Iron + y.Iron,
		Caffeine:           x.Caffeine + y.Caffeine,
		VitaminA:           x.VitaminA + y.VitaminA,
		VitaminB1:          x.VitaminB1 + y.VitaminB1,
		VitaminB2:          x.VitaminB2 + y.VitaminB2,
		VitaminB3:          x.VitaminB3 + y.VitaminB3,
		VitaminB4:          x.VitaminB4 + y.VitaminB4,
		VitaminB5:          x.VitaminB5 + y.VitaminB5,
		VitaminB6:          x.VitaminB6 + y.VitaminB6,
		VitaminB12:         x.VitaminB12 + y.VitaminB12,
		VitaminC:           x.VitaminC + y.VitaminC,
		VitaminD:           x.VitaminD + y.VitaminD,
		VitaminE:           x.VitaminE + y.VitaminE,
		VitaminK:           x.VitaminK + y.VitaminK,
	}
}
