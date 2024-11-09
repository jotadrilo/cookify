package fs

import (
	"context"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/jotadrilo/cookify/app/adapters/repositories/fs/model"
	"github.com/jotadrilo/cookify/app/adapters/repositories/unimpl"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/slices"
)

type RecipesRepository struct {
	unimpl.RecipesRepository

	root string
}

var _ ports.RecipesRepository = (*RecipesRepository)(nil)

type RecipesRepositoryOptions struct {
	Root string
}

func NewRecipesRepository(opts *RecipesRepositoryOptions) *RecipesRepository {
	return &RecipesRepository{
		root: opts.Root,
	}
}

func (x *RecipesRepository) getFile() string {
	return filepath.Join(x.root, "recipes.json")
}

func (x *RecipesRepository) ListRecipes(_ context.Context) ([]*domain.Recipe, error) {
	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map(items, model.RecipeToDomainRecipe), nil
}

func (x *RecipesRepository) GetRecipeByUUID(_ context.Context, recipeID string) (*domain.Recipe, error) {
	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == recipeID {
			return model.RecipeToDomainRecipe(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("recipe")
}

func (x *RecipesRepository) CreateUserRecipe(ctx context.Context, userID string, v *domain.Recipe) (*domain.Recipe, error) {
	if vv, err := x.GetUserRecipeByName(ctx, userID, v.Name); err == nil {
		return vv, errorutils.NewErrAlreadyExists("recipe")
	}

	var (
		vv = model.DomainRecipeToRecipe(v)
	)

	vv.UUID = uuid.New().String()
	vv.UserUUID = userID

	if err := appendToJSON(x.getFile(), vv); err != nil {
		return nil, err
	}

	return model.RecipeToDomainRecipe(vv), nil
}

func (x *RecipesRepository) ListUserRecipes(_ context.Context, userID string) ([]*domain.Recipe, error) {
	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}

	return slices.Map(slices.Select(items, model.User{UUID: userID}.OwnsRecipe), model.RecipeToDomainRecipe), nil
}

func (x *RecipesRepository) GetUserRecipeByUUID(_ context.Context, userID string, recipeID string) (*domain.Recipe, error) {
	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == recipeID && (model.User{UUID: userID}).OwnsRecipe(item) {
			return model.RecipeToDomainRecipe(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("recipe")
}

func (x *RecipesRepository) UpdateUserRecipeByUUID(_ context.Context, userID string, recipeID string, v *domain.Recipe) (*domain.Recipe, error) {
	var vv = model.DomainRecipeToRecipe(v)

	vv.UUID = recipeID
	vv.UserUUID = userID

	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}

	for ix, item := range items {
		if item.UUID == recipeID && (model.User{UUID: userID}).OwnsRecipe(item) {
			items[ix] = vv
			break
		}
	}

	if err := writeJSON(x.getFile(), items); err != nil {
		return nil, err
	}

	return model.RecipeToDomainRecipe(vv), nil
}

func (x *RecipesRepository) GetUserRecipeByName(_ context.Context, userID string, name string) (*domain.Recipe, error) {
	items, err := decodeJSON[*model.Recipe](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Name == name && (model.User{UUID: userID}).OwnsRecipe(item) {
			return model.RecipeToDomainRecipe(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("recipe")
}
