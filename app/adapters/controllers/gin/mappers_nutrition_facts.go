package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
)

func DomainNutritionFactsToAPINutritionFacts(v *domain.NutritionFacts) *api.NutritionFacts {
	if v == nil {
		return nil
	}
	return &api.NutritionFacts{
		Cal:                v.Cal,
		Caffeine:           oapi.Float32(v.Caffeine),
		Calcium:            oapi.Float32(v.Calcium),
		CarbohydrateSugar:  oapi.Float32(v.CarbohydrateSugar),
		CarbohydrateTotal:  oapi.Float32(v.CarbohydrateTotal),
		Cholesterol:        oapi.Float32(v.Cholesterol),
		FatMonounsaturated: oapi.Float32(v.FatMonounsaturated),
		FatPolyunsaturated: oapi.Float32(v.FatPolyunsaturated),
		FatSaturated:       oapi.Float32(v.FatSaturated),
		FatTotal:           oapi.Float32(v.FatTotal),
		Fiber:              oapi.Float32(v.Fiber),
		Iron:               oapi.Float32(v.Iron),
		Potassium:          oapi.Float32(v.Potassium),
		Protein:            oapi.Float32(v.Protein),
		Salt:               oapi.Float32(v.Salt),
		Sodium:             oapi.Float32(v.Sodium),
		VitaminA:           oapi.Float32(v.VitaminA),
		VitaminB1:          oapi.Float32(v.VitaminB1),
		VitaminB12:         oapi.Float32(v.VitaminB12),
		VitaminB2:          oapi.Float32(v.VitaminB2),
		VitaminB3:          oapi.Float32(v.VitaminB3),
		VitaminB4:          oapi.Float32(v.VitaminB4),
		VitaminB5:          oapi.Float32(v.VitaminB5),
		VitaminB6:          oapi.Float32(v.VitaminB6),
		VitaminC:           oapi.Float32(v.VitaminC),
		VitaminD:           oapi.Float32(v.VitaminD),
		VitaminE:           oapi.Float32(v.VitaminE),
		VitaminK:           oapi.Float32(v.VitaminK),
	}
}

func NutritionFactsToDomainNutritionFacts(v *api.NutritionFacts) *domain.NutritionFacts {
	if v == nil {
		return nil
	}
	return &domain.NutritionFacts{
		Cal:                v.Cal,
		FatTotal:           oapi.Float32Value(v.FatTotal),
		FatSaturated:       oapi.Float32Value(v.FatSaturated),
		FatMonounsaturated: oapi.Float32Value(v.FatMonounsaturated),
		FatPolyunsaturated: oapi.Float32Value(v.FatPolyunsaturated),
		Cholesterol:        oapi.Float32Value(v.Cholesterol),
		Salt:               oapi.Float32Value(v.Salt),
		Sodium:             oapi.Float32Value(v.Sodium),
		Potassium:          oapi.Float32Value(v.Potassium),
		CarbohydrateTotal:  oapi.Float32Value(v.CarbohydrateTotal),
		CarbohydrateSugar:  oapi.Float32Value(v.CarbohydrateSugar),
		Protein:            oapi.Float32Value(v.Protein),
		Fiber:              oapi.Float32Value(v.Fiber),
		Calcium:            oapi.Float32Value(v.Calcium),
		Iron:               oapi.Float32Value(v.Iron),
		Caffeine:           oapi.Float32Value(v.Caffeine),
		VitaminA:           oapi.Float32Value(v.VitaminA),
		VitaminB1:          oapi.Float32Value(v.VitaminB1),
		VitaminB2:          oapi.Float32Value(v.VitaminB2),
		VitaminB3:          oapi.Float32Value(v.VitaminB3),
		VitaminB4:          oapi.Float32Value(v.VitaminB4),
		VitaminB5:          oapi.Float32Value(v.VitaminB5),
		VitaminB6:          oapi.Float32Value(v.VitaminB6),
		VitaminB12:         oapi.Float32Value(v.VitaminB12),
		VitaminC:           oapi.Float32Value(v.VitaminC),
		VitaminD:           oapi.Float32Value(v.VitaminD),
		VitaminE:           oapi.Float32Value(v.VitaminE),
		VitaminK:           oapi.Float32Value(v.VitaminK),
	}
}
