package model

import "github.com/jotadrilo/cookify/app/core/domain"

type NutritionFacts struct {
	UUID               string  `json:"uuid"`
	Cal                float32 `json:"cal"`
	FatTotal           float32 `json:"fat_total,omitempty"`
	FatSaturated       float32 `json:"fat_saturated,omitempty"`
	FatMonounsaturated float32 `json:"fat_monounsaturated,omitempty"`
	FatPolyunsaturated float32 `json:"fat_polyunsaturated,omitempty"`
	Cholesterol        float32 `json:"cholesterol,omitempty"`
	Salt               float32 `json:"salt,omitempty"`
	Sodium             float32 `json:"sodium,omitempty"`
	Potassium          float32 `json:"potassium,omitempty"`
	CarbohydrateTotal  float32 `json:"carbohydrate_total,omitempty"`
	CarbohydrateSugar  float32 `json:"carbohydrate_sugar,omitempty"`
	Protein            float32 `json:"protein,omitempty"`
	Fiber              float32 `json:"fiber,omitempty"`
	Calcium            float32 `json:"calcium,omitempty"`
	Iron               float32 `json:"iron,omitempty"`
	Caffeine           float32 `json:"caffeine,omitempty"`
	VitaminA           float32 `json:"vitamin_a,omitempty"`
	VitaminB1          float32 `json:"vitamin_b1,omitempty"`
	VitaminB2          float32 `json:"vitamin_b2,omitempty"`
	VitaminB3          float32 `json:"vitamin_b3,omitempty"`
	VitaminB4          float32 `json:"vitamin_b4,omitempty"`
	VitaminB5          float32 `json:"vitamin_b5,omitempty"`
	VitaminB6          float32 `json:"vitamin_b6,omitempty"`
	VitaminB12         float32 `json:"vitamin_b12,omitempty"`
	VitaminC           float32 `json:"vitamin_c,omitempty"`
	VitaminD           float32 `json:"vitamin_d,omitempty"`
	VitaminE           float32 `json:"vitamin_e,omitempty"`
	VitaminK           float32 `json:"vitamin_k,omitempty"`
}

func NutritionFactsToDomainNutritionFacts(x *NutritionFacts) *domain.NutritionFacts {
	if x == nil {
		return nil
	}

	return &domain.NutritionFacts{
		UUID:               x.UUID,
		Cal:                x.Cal,
		FatTotal:           x.FatTotal,
		FatSaturated:       x.FatSaturated,
		FatMonounsaturated: x.FatMonounsaturated,
		FatPolyunsaturated: x.FatPolyunsaturated,
		Cholesterol:        x.Cholesterol,
		Salt:               x.Salt,
		Sodium:             x.Sodium,
		Potassium:          x.Potassium,
		CarbohydrateTotal:  x.CarbohydrateTotal,
		CarbohydrateSugar:  x.CarbohydrateSugar,
		Protein:            x.Protein,
		Fiber:              x.Fiber,
		Calcium:            x.Calcium,
		Iron:               x.Iron,
		Caffeine:           x.Caffeine,
		VitaminA:           x.VitaminA,
		VitaminB1:          x.VitaminB1,
		VitaminB2:          x.VitaminB2,
		VitaminB3:          x.VitaminB3,
		VitaminB4:          x.VitaminB4,
		VitaminB5:          x.VitaminB5,
		VitaminB6:          x.VitaminB6,
		VitaminB12:         x.VitaminB12,
		VitaminC:           x.VitaminC,
		VitaminD:           x.VitaminD,
		VitaminE:           x.VitaminE,
		VitaminK:           x.VitaminK,
	}
}

func DomainNutritionFactsToNutritionFacts(x *domain.NutritionFacts) *NutritionFacts {
	if x == nil {
		return nil
	}

	var facts = &NutritionFacts{
		UUID:               x.UUID,
		Cal:                x.Cal,
		FatTotal:           x.FatTotal,
		FatSaturated:       x.FatSaturated,
		FatMonounsaturated: x.FatMonounsaturated,
		FatPolyunsaturated: x.FatPolyunsaturated,
		Cholesterol:        x.Cholesterol,
		Salt:               x.Salt,
		Sodium:             x.Sodium,
		Potassium:          x.Potassium,
		CarbohydrateTotal:  x.CarbohydrateTotal,
		CarbohydrateSugar:  x.CarbohydrateSugar,
		Protein:            x.Protein,
		Fiber:              x.Fiber,
		Calcium:            x.Calcium,
		Iron:               x.Iron,
		Caffeine:           x.Caffeine,
		VitaminA:           x.VitaminA,
		VitaminB1:          x.VitaminB1,
		VitaminB2:          x.VitaminB2,
		VitaminB3:          x.VitaminB3,
		VitaminB4:          x.VitaminB4,
		VitaminB5:          x.VitaminB5,
		VitaminB6:          x.VitaminB6,
		VitaminB12:         x.VitaminB12,
		VitaminC:           x.VitaminE,
		VitaminD:           x.VitaminD,
		VitaminE:           x.VitaminK,
		VitaminK:           x.VitaminK,
	}

	return facts
}
