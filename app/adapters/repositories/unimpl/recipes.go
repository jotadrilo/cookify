package unimpl

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type RecipesRepository struct {
	ports.Repository
}

var _ ports.RecipesRepository = (*RecipesRepository)(nil)

func (x *RecipesRepository) ListRecipes(context.Context) ([]*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("ListRecipes")
}

func (x *RecipesRepository) GetRecipeByUUID(context.Context, string) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("GetRecipeByUUID")
}

func (x *RecipesRepository) CreateUserRecipe(context.Context, string, *domain.Recipe) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserRecipe")
}

func (x *RecipesRepository) ListUserRecipes(context.Context, string) ([]*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserRecipes")
}

func (x *RecipesRepository) GetUserRecipeByUUID(context.Context, string, string) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserRecipeByUUID")
}

func (x *RecipesRepository) UpdateUserRecipeByUUID(context.Context, string, string, *domain.Recipe) (*domain.Recipe, error) {
	return nil, errorutils.NewErrNotImplemented("UpdateUserRecipeByUUID")
}
