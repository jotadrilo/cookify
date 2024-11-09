package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type Recipe struct {
	bun.BaseModel `bun:"table:recipes"`

	ID int64 `bun:"id,pk,autoincrement"`

	UUID        string        `bun:"uuid,pk"`
	Name        string        `bun:"name"`
	Ingredients []*Ingredient `bun:"rel:has-many,join:uuid=recipe_uuid"`
}

var _ bun.BeforeAppendModelHook = (*Recipe)(nil)

func (r *Recipe) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.UUID == "" {
		r.UUID = uuid.NewString()
	}

	return nil
}

func RecipeToDomainRecipe(x *Recipe) *domain.Recipe {
	if x == nil {
		return nil
	}

	var ingredients []*domain.Ingredient

	for _, i := range x.Ingredients {
		ingredients = append(ingredients, IngredientToDomainIngredient(i))
	}

	var r = &domain.Recipe{
		UUID:        x.UUID,
		Name:        x.Name,
		Ingredients: ingredients,
	}

	r.Fixup()

	return r
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
		Name:        x.Name,
		Ingredients: ingredients,
	}
}
