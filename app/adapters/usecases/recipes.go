package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
)

type UnimplementedRecipesUseCase struct {
	ports.UseCase
}

var _ ports.RecipesUseCase = (*UnimplementedRecipesUseCase)(nil)

func (x *UnimplementedRecipesUseCase) ListRecipes(context.Context) ([]*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("ListRecipes")
}

func (x *UnimplementedRecipesUseCase) GetRecipeByUUID(context.Context, string) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("GetRecipeByUUID")
}

func (x *UnimplementedRecipesUseCase) CreateUserRecipe(context.Context, string, *domain.Recipe) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserRecipe")
}

func (x *UnimplementedRecipesUseCase) ListUserRecipes(context.Context, string) ([]*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserRecipes")
}

func (x *UnimplementedRecipesUseCase) GetUserRecipeByUUID(context.Context, string, string) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserRecipeByUUID")
}

func (x *UnimplementedRecipesUseCase) UpdateUserRecipeByUUID(context.Context, string, string, *domain.Recipe) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("UpdateUserRecipeByUUID")
}

type RecipesUseCase struct {
	UnimplementedRecipesUseCase

	Recipes ports.RecipesRepository
}

type RecipesUseCaseOptions struct {
	Recipes ports.RecipesRepository
}

func NewRecipesUseCase(opts *RecipesUseCaseOptions) *RecipesUseCase {
	return &RecipesUseCase{
		Recipes: opts.Recipes,
	}
}

func (x *RecipesUseCase) ListRecipes(ctx context.Context) ([]*domain.Recipe, error) {
	return x.Recipes.ListRecipes(ctx)
}

func (x *RecipesUseCase) GetRecipeByUUID(ctx context.Context, recipeID string) (*domain.Recipe, error) {
	return x.Recipes.GetRecipeByUUID(ctx, recipeID)
}

func (x *RecipesUseCase) CreateUserRecipe(ctx context.Context, userID string, v *domain.Recipe) (*domain.Recipe, error) {
	vv, err := x.Recipes.CreateUserRecipe(ctx, userID, v)
	if err != nil {
		logger.Errorf("Cannot create recipe %v: %s", v, err.Error())

		if errors.Is(err, errorutils.ErrAlreadyExists) {
			return nil, errorutils.NewErrAlreadyExists(fmt.Sprintf("recipe %s", vv.UUID))
		}

		return nil, errorutils.NewErrNotCreated("recipe")
	}

	logger.Infof("Created recipe %q", vv.UUID)

	return vv, nil
}

func (x *RecipesUseCase) ListUserRecipes(ctx context.Context, userID string) ([]*domain.Recipe, error) {
	return x.Recipes.ListUserRecipes(ctx, userID)
}

func (x *RecipesUseCase) GetUserRecipeByUUID(ctx context.Context, userID string, recipeID string) (*domain.Recipe, error) {
	return x.Recipes.GetUserRecipeByUUID(ctx, userID, recipeID)
}

func (x *RecipesUseCase) UpdateUserRecipeByUUID(ctx context.Context, userID string, recipeID string, v *domain.Recipe) (*domain.Recipe, error) {
	return x.Recipes.UpdateUserRecipeByUUID(ctx, userID, recipeID, v)
}
