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

type MenusRepository struct {
	unimpl.MenusRepository

	db *bun.DB
}

type MenusRepositoryOptions struct {
	DB *bun.DB
}

func NewMenusRepository(opts *MenusRepositoryOptions) *MenusRepository {
	return &MenusRepository{
		db: opts.DB,
	}
}

func (x *MenusRepository) CreateMenu(ctx context.Context, v *domain.Menu) (*domain.Menu, error) {
	if pp, err := x.GetMenuByName(ctx, v.Name); err == nil {
		return pp, nil
	}

	var menu = model.DomainMenuToMenu(v)

	// Use a transaction for atomicity
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// Insert the Menu
		if _, err := tx.NewInsert().
			Model(menu).
			Exec(ctx); err != nil {
			return err
		}

		// Insert Ingredients (set MenuUUID to match the Menu's UUID)
		for _, ingredient := range menu.Ingredients {
			ingredient.MenuUUID = menu.UUID
		}

		if len(menu.Ingredients) > 0 {
			if _, err := tx.NewInsert().
				Model(&menu.Ingredients).
				Exec(ctx); err != nil {
				return err
			}
		}

		for _, recipe := range menu.Recipes {
			if _, err := tx.NewInsert().
				Model(&model.MenuRecipe{
					MenuUUID:   menu.UUID,
					RecipeUUID: recipe.UUID,
				}).
				Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return model.MenuToDomainMenu(menu), nil
}

func (x *MenusRepository) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	var (
		items []*domain.Menu
		menus []*model.Menu
	)

	err := x.db.NewSelect().
		Model(&menus).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range menus {
		items = append(items, model.MenuToDomainMenu(p))
	}

	return items, nil
}

func (x *MenusRepository) GetMenuByUUID(ctx context.Context, uuid string) (*domain.Menu, error) {
	var menu model.Menu

	err := x.db.NewSelect().
		Model(&menu).
		Where("uuid = ?", uuid).
		// Load Recipes > Ingredients > Product > NutritionFacts
		Relation("Recipes.Ingredients.Product.NutritionFacts").
		// Load Ingredients > Product > NutritionFacts
		Relation("Ingredients.Product.NutritionFacts").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("menu")
		}
		return nil, err
	}

	return model.MenuToDomainMenu(&menu), nil
}

func (x *MenusRepository) GetMenuByName(ctx context.Context, name string) (*domain.Menu, error) {
	var menu model.Menu

	err := x.db.NewSelect().
		Model(&menu).
		Where("name = ?", name).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("menu")
		}
		return nil, err
	}

	return model.MenuToDomainMenu(&menu), nil
}
