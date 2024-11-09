package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
)

func DomainRecipeToAPIRecipe(v *domain.Recipe) *api.Recipe {
	if v == nil {
		return nil
	}

	return &api.Recipe{
		Uuid:                oapi.UUID(v.UUID),
		Name:                v.Name,
		Ingredients:         DomainIngredientsToAPIIngredients(v.Ingredients),
		NutritionFacts100:   DomainNutritionFactsToAPINutritionFacts(v.NutritionFacts),
		NutritionFactsTotal: DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal),
		Quantity:            oapi.Float32(v.Quantity),
	}
}

func DomainRecipesToAPIRecipes(s []*domain.Recipe) []api.Recipe {
	var recipes []api.Recipe

	for _, recipe := range s {
		recipes = append(recipes, *DomainRecipeToAPIRecipe(recipe))
	}

	return recipes
}

func RecipeToDomainRecipe(v *api.Recipe) *domain.Recipe {
	if v == nil {
		return nil
	}

	return &domain.Recipe{
		UUID:                oapi.UUIDValue(v.Uuid),
		Name:                v.Name,
		Ingredients:         IngredientsToDomainIngredients(v.Ingredients),
		Quantity:            oapi.Float32Value(v.Quantity),
		NutritionFacts:      NutritionFactsToDomainNutritionFacts(v.NutritionFacts100),
		NutritionFactsTotal: NutritionFactsToDomainNutritionFacts(v.NutritionFactsTotal),
	}
}

func RecipePostRequestToDomainRecipe(v *api.PostUsersParamUserIDRecipesJSONRequestBody) *domain.Recipe {
	if v == nil {
		return nil
	}

	return &domain.Recipe{
		Name:                v.Name,
		Ingredients:         IngredientsToDomainIngredients(v.Ingredients),
		Quantity:            oapi.Float32Value(v.Quantity),
		NutritionFacts:      NutritionFactsToDomainNutritionFacts(v.NutritionFacts100),
		NutritionFactsTotal: NutritionFactsToDomainNutritionFacts(v.NutritionFactsTotal),
	}
}

func RecipePatchRequestToDomainRecipe(v *api.PatchUsersParamUserIDRecipesParamRecipeIDJSONRequestBody) *domain.Recipe {
	if v == nil {
		return nil
	}

	var recipe domain.Recipe

	if v.Ingredients != nil {
		recipe.Ingredients = IngredientsToDomainIngredients(*v.Ingredients)
	}

	if v.Name != nil {
		recipe.Name = *v.Name
	}

	return &recipe
}
