package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/adapters/repositories/pg/model"
	"github.com/jotadrilo/cookify/app/adapters/repositories/unimpl"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type RecipesRepository struct {
	unimpl.RecipesRepository

	db *bun.DB
}

type RecipesRepositoryOptions struct {
	DB *bun.DB
}

func NewRecipesRepository(opts *RecipesRepositoryOptions) *RecipesRepository {
	return &RecipesRepository{
		db: opts.DB,
	}
}

func (x *RecipesRepository) CreateRecipe(ctx context.Context, v *domain.Recipe) (*domain.Recipe, error) {
	if pp, err := x.GetRecipeByName(ctx, v.Name); err == nil {
		return pp, nil
	}

	var recipe = model.DomainRecipeToRecipe(v)

	// Use a transaction for atomicity
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// Insert the Recipe
		if _, err := tx.NewInsert().
			Model(recipe).
			Exec(ctx); err != nil {
			return err
		}

		// Insert Ingredients (set RecipeUUID to match the Recipe's UUID)
		for _, ingredient := range recipe.Ingredients {
			ingredient.RecipeUUID = recipe.UUID
		}

		if _, err := tx.NewInsert().
			Model(&recipe.Ingredients).
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return model.RecipeToDomainRecipe(recipe), nil
}

func (x *RecipesRepository) ListRecipes(ctx context.Context) ([]*domain.Recipe, error) {
	var (
		items   []*domain.Recipe
		recipes []*model.Recipe
	)

	err := x.db.NewSelect().
		Model(&recipes).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range recipes {
		items = append(items, model.RecipeToDomainRecipe(p))
	}

	return items, nil
}

func (x *RecipesRepository) GetRecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	var recipe model.Recipe

	err := x.db.NewSelect().
		Model(&recipe).
		Where("uuid = ?", uuid).
		// Load Ingredients > Product > NutritionFacts
		Relation("Ingredients.Product.NutritionFacts").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("recipe")
		}
		return nil, err
	}

	return model.RecipeToDomainRecipe(&recipe), nil
}

func (x *RecipesRepository) GetRecipeByName(ctx context.Context, name string) (*domain.Recipe, error) {
	var recipe model.Recipe

	err := x.db.NewSelect().
		Model(&recipe).
		Where("name = ?", name).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("recipe")
		}
		return nil, err
	}

	return model.RecipeToDomainRecipe(&recipe), nil
}
