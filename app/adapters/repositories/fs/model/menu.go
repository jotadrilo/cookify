package model

import "github.com/jotadrilo/cookify/app/core/domain"

type Menu struct {
	UUID        string        `json:"uuid"`
	UserUUID    string        `json:"user_uuid"`
	Name        string        `json:"name"`
	Label       string        `json:"label"`
	Recipes     []*Recipe     `json:"recipes"`
	Ingredients []*Ingredient `json:"ingredients"`
}

func MenuToDomainMenu(x *Menu) *domain.Menu {
	if x == nil {
		return nil
	}

	var (
		recipes     []*domain.Recipe
		ingredients []*domain.Ingredient
	)

	for _, recipe := range x.Recipes {
		recipes = append(recipes, RecipeToDomainRecipe(recipe))
	}

	for _, ingredient := range x.Ingredients {
		ingredients = append(ingredients, IngredientToDomainIngredient(ingredient))
	}

	var menu = &domain.Menu{
		UUID:        x.UUID,
		UserUUID:    x.UserUUID,
		Name:        x.Name,
		Label:       domain.ParseMenuLabel(x.Label),
		Recipes:     recipes,
		Ingredients: ingredients,
	}

	menu.Fixup()

	return menu
}

func DomainMenuToMenu(x *domain.Menu) *Menu {
	if x == nil {
		return nil
	}

	var (
		recipes     []*Recipe
		ingredients []*Ingredient
	)

	for _, recipe := range x.Recipes {
		recipes = append(recipes, DomainRecipeToRecipe(recipe))
	}

	for _, ingredient := range x.Ingredients {
		ingredients = append(ingredients, DomainIngredientToIngredient(ingredient))
	}

	return &Menu{
		UUID:        x.UUID,
		UserUUID:    x.UserUUID,
		Name:        x.Name,
		Label:       x.Label.String(),
		Recipes:     recipes,
		Ingredients: ingredients,
	}
}
