package domain

import (
	"sort"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	UUID                string
	Slug                string
	DisplayNameLangEsES string
	DisplayNameLangEnUS string
	Vendors             []string
	CurrentPrice        *Price
	Unit                Unit
	NutritionFacts      *NutritionFacts
}

func (a *Product) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Slug, validation.Required),
		validation.Field(&a.DisplayNameLangEsES, validation.Required),
		validation.Field(&a.Unit, validation.Required),
	)
}

// FindTopEquivalentProducts finds the top n most similar products from a list and assigns a similarity score
func (x *Product) FindTopEquivalentProducts(n int, threshold float32, options ...*Product) []*EquivalentProduct {
	if x == nil || x.NutritionFacts == nil || len(options) == 0 {
		return nil
	}

	// Calculate distances and sort them by score in descending order

	type distance struct {
		Index    int
		Distance float32
		Score    float32
	}

	var (
		distances  = make([]distance, len(options))
		xFactsNorm = x.NutritionFacts.Norm()
	)

	for i, option := range options {
		var (
			optFactsNorm = option.NutritionFacts.Norm()
			dist         = xFactsNorm.EuclideanDistance(optFactsNorm)
		)

		distances[i] = distance{
			Index:    i,
			Distance: dist,
			Score:    1 - dist,
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Score > distances[j].Score
	})

	topSimilar := make([]*EquivalentProduct, 0, n)

	for i := 0; i < n && i < len(distances); i++ {
		var (
			dist = distances[i]
			y    = options[dist.Index]
		)

		if dist.Score < threshold {
			continue
		}

		diff := &NutritionFacts{
			Cal:                x.NutritionFacts.Cal - y.NutritionFacts.Cal,
			FatTotal:           x.NutritionFacts.FatTotal - y.NutritionFacts.FatTotal,
			FatSaturated:       x.NutritionFacts.FatSaturated - y.NutritionFacts.FatSaturated,
			FatMonounsaturated: x.NutritionFacts.FatMonounsaturated - y.NutritionFacts.FatMonounsaturated,
			FatPolyunsaturated: x.NutritionFacts.FatPolyunsaturated - y.NutritionFacts.FatPolyunsaturated,
			Cholesterol:        x.NutritionFacts.Cholesterol - y.NutritionFacts.Cholesterol,
			Salt:               x.NutritionFacts.Salt - y.NutritionFacts.Salt,
			Sodium:             x.NutritionFacts.Sodium - y.NutritionFacts.Sodium,
			Potassium:          x.NutritionFacts.Potassium - y.NutritionFacts.Potassium,
			CarbohydrateTotal:  x.NutritionFacts.CarbohydrateTotal - y.NutritionFacts.CarbohydrateTotal,
			CarbohydrateSugar:  x.NutritionFacts.CarbohydrateSugar - y.NutritionFacts.CarbohydrateSugar,
			Protein:            x.NutritionFacts.Protein - y.NutritionFacts.Protein,
			Fiber:              x.NutritionFacts.Fiber - y.NutritionFacts.Fiber,
			Calcium:            x.NutritionFacts.Calcium - y.NutritionFacts.Calcium,
			Iron:               x.NutritionFacts.Iron - y.NutritionFacts.Iron,
			Caffeine:           x.NutritionFacts.Caffeine - y.NutritionFacts.Caffeine,
			VitaminA:           x.NutritionFacts.VitaminA - y.NutritionFacts.VitaminA,
			VitaminB1:          x.NutritionFacts.VitaminB1 - y.NutritionFacts.VitaminB1,
			VitaminB2:          x.NutritionFacts.VitaminB2 - y.NutritionFacts.VitaminB2,
			VitaminB3:          x.NutritionFacts.VitaminB3 - y.NutritionFacts.VitaminB3,
			VitaminB4:          x.NutritionFacts.VitaminB4 - y.NutritionFacts.VitaminB4,
			VitaminB5:          x.NutritionFacts.VitaminB5 - y.NutritionFacts.VitaminB5,
			VitaminB6:          x.NutritionFacts.VitaminB6 - y.NutritionFacts.VitaminB6,
			VitaminB12:         x.NutritionFacts.VitaminB12 - y.NutritionFacts.VitaminB12,
			VitaminC:           x.NutritionFacts.VitaminC - y.NutritionFacts.VitaminC,
			VitaminD:           x.NutritionFacts.VitaminD - y.NutritionFacts.VitaminD,
			VitaminE:           x.NutritionFacts.VitaminE - y.NutritionFacts.VitaminE,
			VitaminK:           x.NutritionFacts.VitaminK - y.NutritionFacts.VitaminK,
		}

		topSimilar = append(topSimilar, &EquivalentProduct{
			Score:              dist.Score,
			Product:            y,
			NutritionFactsDiff: diff,
		})
	}

	return topSimilar
}
