package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type Menu struct {
	bun.BaseModel `bun:"table:menus"`

	ID int64 `bun:"id,pk,autoincrement"`

	UUID        string        `bun:"uuid,pk"`
	Name        string        `bun:"name"`
	Label       string        `bun:"label"`
	Recipes     []*Recipe     `bun:"m2m:menu_recipes,join:Menu=Recipe"`
	Ingredients []*Ingredient `bun:"rel:has-many,join:uuid=menu_uuid"`
}

type MenuRecipe struct {
	ID int64 `bun:"id,pk,autoincrement"`

	MenuUUID   string  `bun:"menu_uuid,notnull"`
	Menu       *Menu   `bun:"rel:belongs-to,join:menu_uuid=uuid"`
	RecipeUUID string  `bun:"recipe_uuid,notnull"`
	Recipe     *Recipe `bun:"rel:belongs-to,join:recipe_uuid=uuid"`
}

var _ bun.BeforeAppendModelHook = (*Menu)(nil)

func (r *Menu) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.UUID == "" {
		r.UUID = uuid.NewString()
	}

	return nil
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
		Name:        x.Name,
		Label:       x.Label.String(),
		Recipes:     recipes,
		Ingredients: ingredients,
	}
}
