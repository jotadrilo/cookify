package model

import "github.com/jotadrilo/cookify/app/core/domain"

type Recipe struct {
	UUID        string        `json:"uuid"`
	UserUUID    string        `json:"user_uuid"`
	Name        string        `json:"name"`
	Ingredients []*Ingredient `json:"ingredients,omitempty"`
}

func RecipeToDomainRecipe(x *Recipe) *domain.Recipe {
	if x == nil {
		return nil
	}

	var ingredients []*domain.Ingredient

	for _, i := range x.Ingredients {
		ingredients = append(ingredients, IngredientToDomainIngredient(i))
	}

	var recipe = &domain.Recipe{
		UUID:        x.UUID,
		UserUUID:    x.UserUUID,
		Name:        x.Name,
		Ingredients: ingredients,
	}

	recipe.Fixup()

	return recipe
}

func DomainRecipeToRecipe(x *domain.Recipe) *Recipe {
	if x == nil {
		return nil
	}

	var ingredients []*Ingredient

	for _, i := range x.Ingredients {
		ingredients = append(ingredients, DomainIngredientToIngredient(i))
	}

	return &Recipe{
		UUID:        x.UUID,
		UserUUID:    x.UserUUID,
		Name:        x.Name,
		Ingredients: ingredients,
	}
}
