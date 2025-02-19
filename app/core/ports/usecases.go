package ports

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
)

// UseCase is a marker to ensure that all UseCase interfaces are unique
type UseCase interface {
	useCase()
}

type UsersUseCase interface {
	UseCase

	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]*domain.User, error)
	GetUserByUUID(ctx context.Context, userID string) (*domain.User, error)
}

type ProductsUseCase interface {
	UseCase

	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	ListProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductByUUID(ctx context.Context, productID string) (*domain.Product, error)
	GetProductEquivalentsByUUID(ctx context.Context, productID string) ([]*domain.EquivalentProduct, error)
}

type RecipesUseCase interface {
	UseCase

	ListRecipes(ctx context.Context) ([]*domain.Recipe, error)
	GetRecipeByUUID(ctx context.Context, recipeID string) (*domain.Recipe, error)
	CreateUserRecipe(ctx context.Context, userID string, recipe *domain.Recipe) (*domain.Recipe, error)
	ListUserRecipes(ctx context.Context, userID string) ([]*domain.Recipe, error)
	GetUserRecipeByUUID(ctx context.Context, userID string, recipeID string) (*domain.Recipe, error)
	UpdateUserRecipeByUUID(ctx context.Context, userID string, recipeID string, recipe *domain.Recipe) (*domain.Recipe, error)
}

type MenusUseCase interface {
	UseCase

	ListMenus(ctx context.Context) ([]*domain.Menu, error)
	GetMenuByUUID(ctx context.Context, menuID string) (*domain.Menu, error)
	CreateUserMenu(ctx context.Context, userID string, menu *domain.Menu) (*domain.Menu, error)
	ListUserMenus(ctx context.Context, userID string) ([]*domain.Menu, error)
	GetUserMenuByUUID(ctx context.Context, userID string, menuID string) (*domain.Menu, error)
}

type DailyMenusUseCase interface {
	UseCase

	ListDailyMenus(ctx context.Context) ([]*domain.DailyMenu, error)
	GetDailyMenuByUUID(ctx context.Context, dailyMenuID string) (*domain.DailyMenu, error)
	CreateUserDailyMenu(ctx context.Context, userID string, dailyMenu *domain.DailyMenu) (*domain.DailyMenu, error)
	ListUserDailyMenus(ctx context.Context, userID string) ([]*domain.DailyMenu, error)
	GetUserDailyMenuByUUID(ctx context.Context, userID string, dailyMenuID string) (*domain.DailyMenu, error)
}
