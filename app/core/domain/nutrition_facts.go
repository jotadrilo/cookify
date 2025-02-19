package domain

import (
	"math"

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
	if x == nil {
		return nil
	}

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

func (x *NutritionFacts) Norm() *NutritionFacts {
	if x == nil {
		return nil
	}

	// Calculate the L2 norm (magnitude) of the vector
	// It only takes most important nutrition facts into account

	var norm float32

	for _, val := range []float32{
		//x.Cal,
		x.FatTotal,
		x.FatSaturated,
		x.FatMonounsaturated,
		x.FatPolyunsaturated,
		//x.Cholesterol,
		//x.Salt,
		//x.Sodium,
		//x.Potassium,
		x.CarbohydrateTotal,
		x.CarbohydrateSugar,
		x.Protein,
		x.Fiber,
		//x.Calcium,
		//x.Iron,
		//x.Caffeine,
		//x.VitaminA,
		//x.VitaminB1,
		//x.VitaminB2,
		//x.VitaminB3,
		//x.VitaminB4,
		//x.VitaminB5,
		//x.VitaminB6,
		//x.VitaminB12,
		//x.VitaminC,
		//x.VitaminD,
		//x.VitaminE,
		//x.VitaminK,
	} {
		norm += val * val
	}

	norm = float32(math.Sqrt(float64(norm)))

	// Normalize each field by dividing by the L2 norm
	return x.Multiply(1 / norm)
}

// EuclideanDistance calculates the Euclidean distance between two NutritionFacts structs
func (x *NutritionFacts) EuclideanDistance(y *NutritionFacts) float32 {
	if x == nil || y == nil {
		return 1
	}

	return float32(math.Sqrt(
		//math.Pow(float64(x.Cal-y.Cal), 2) +
		math.Pow(float64(x.FatTotal-y.FatTotal), 2) +
			math.Pow(float64(x.FatSaturated-y.FatSaturated), 2) +
			math.Pow(float64(x.FatMonounsaturated-y.FatMonounsaturated), 2) +
			math.Pow(float64(x.FatPolyunsaturated-y.FatPolyunsaturated), 2) +
			//math.Pow(float64(x.Cholesterol-y.Cholesterol), 2) +
			//math.Pow(float64(x.Salt-y.Salt), 2) +
			//math.Pow(float64(x.Sodium-y.Sodium), 2) +
			//math.Pow(float64(x.Potassium-y.Potassium), 2) +
			math.Pow(float64(x.CarbohydrateTotal-y.CarbohydrateTotal), 2) +
			math.Pow(float64(x.CarbohydrateSugar-y.CarbohydrateSugar), 2) +
			math.Pow(float64(x.Protein-y.Protein), 2) +
			math.Pow(float64(x.Fiber-y.Fiber), 2),
		//math.Pow(float64(x.Calcium-y.Calcium), 2) +
		//math.Pow(float64(x.Iron-y.Iron), 2) +
		//math.Pow(float64(x.Caffeine-y.Caffeine), 2) +
		//math.Pow(float64(x.VitaminA-y.VitaminA), 2) +
		//math.Pow(float64(x.VitaminB1-y.VitaminB1), 2) +
		//math.Pow(float64(x.VitaminB2-y.VitaminB2), 2) +
		//math.Pow(float64(x.VitaminB3-y.VitaminB3), 2) +
		//math.Pow(float64(x.VitaminB4-y.VitaminB4), 2) +
		//math.Pow(float64(x.VitaminB5-y.VitaminB5), 2) +
		//math.Pow(float64(x.VitaminB6-y.VitaminB6), 2) +
		//math.Pow(float64(x.VitaminB12-y.VitaminB12), 2) +
		//math.Pow(float64(x.VitaminC-y.VitaminC), 2) +
		//math.Pow(float64(x.VitaminD-y.VitaminD), 2) +
		//math.Pow(float64(x.VitaminE-y.VitaminE), 2) +
		//math.Pow(float64(x.VitaminK-y.VitaminK), 2),
	))
}
