package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type NutritionFacts struct {
	bun.BaseModel `bun:"table:nutrition_facts"`

	ID          int64  `bun:"id,pk,autoincrement"`
	ProductUUID string `bun:"product_uuid,pk"`

	UUID               string  `bun:"uuid,pk"`
	Cal                float32 `bun:"cal"`
	FatTotal           float32 `bun:"fat_total"`
	FatSaturated       float32 `bun:"fat_saturated"`
	FatMonounsaturated float32 `bun:"fat_monounsaturated"`
	FatPolyunsaturated float32 `bun:"fat_polyunsaturated"`
	Cholesterol        float32 `bun:"cholesterol"`
	Salt               float32 `bun:"salt"`
	Sodium             float32 `bun:"sodium"`
	Potassium          float32 `bun:"potassium"`
	CarbohydrateTotal  float32 `bun:"carbohydrate_total"`
	CarbohydrateSugar  float32 `bun:"carbohydrate_sugar"`
	Protein            float32 `bun:"protein"`
	Fiber              float32 `bun:"fiber"`
	Calcium            float32 `bun:"calcium"`
	Iron               float32 `bun:"iron"`
	Caffeine           float32 `bun:"caffeine"`
	VitaminA           float32 `bun:"vitamin_a"`
	VitaminB1          float32 `bun:"vitamin_b1"`
	VitaminB2          float32 `bun:"vitamin_b2"`
	VitaminB3          float32 `bun:"vitamin_b3"`
	VitaminB4          float32 `bun:"vitamin_b4"`
	VitaminB5          float32 `bun:"vitamin_b5"`
	VitaminB6          float32 `bun:"vitamin_b6"`
	VitaminB12         float32 `bun:"vitamin_b12"`
	VitaminC           float32 `bun:"vitamin_c"`
	VitaminD           float32 `bun:"vitamin_d"`
	VitaminE           float32 `bun:"vitamin_e"`
	VitaminK           float32 `bun:"vitamin_k"`
}

var _ bun.BeforeAppendModelHook = (*NutritionFacts)(nil)

func (x *NutritionFacts) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if x.UUID == "" {
		x.UUID = uuid.NewString()
	}

	return nil
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
