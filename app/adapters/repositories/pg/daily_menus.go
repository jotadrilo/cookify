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

type DailyMenusRepository struct {
	unimpl.DailyMenusRepository

	db *bun.DB
}

type DailyMenusRepositoryOptions struct {
	DB *bun.DB
}

func NewDailyMenusRepository(opts *DailyMenusRepositoryOptions) *DailyMenusRepository {
	return &DailyMenusRepository{
		db: opts.DB,
	}
}

func (x *DailyMenusRepository) CreateDailyMenu(ctx context.Context, v *domain.DailyMenu) (*domain.DailyMenu, error) {
	if pp, err := x.GetDailyMenuByName(ctx, v.Name); err == nil {
		return pp, nil
	}

	var dailyMenu = model.DomainDailyMenuToDailyMenu(v)

	// Use a transaction for atomicity
	if err := x.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// Insert the DailyMenu
		if _, err := tx.NewInsert().
			Model(dailyMenu).
			Exec(ctx); err != nil {
			return err
		}

		for _, menu := range dailyMenu.Menus {
			if _, err := tx.NewInsert().
				Model(&model.DailyMenuMenu{
					DailyMenuUUID: dailyMenu.UUID,
					MenuUUID:      menu.UUID,
				}).
				Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return model.DailyMenuToDomainDailyMenu(dailyMenu), nil
}

func (x *DailyMenusRepository) ListDailyMenus(ctx context.Context) ([]*domain.DailyMenu, error) {
	var (
		items      []*domain.DailyMenu
		dailyMenus []*model.DailyMenu
	)

	err := x.db.NewSelect().
		Model(&dailyMenus).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range dailyMenus {
		items = append(items, model.DailyMenuToDomainDailyMenu(p))
	}

	return items, nil
}

func (x *DailyMenusRepository) GetDailyMenuByUUID(ctx context.Context, uuid string) (*domain.DailyMenu, error) {
	var dailyMenu model.DailyMenu

	err := x.db.NewSelect().
		Model(&dailyMenu).
		Where("uuid = ?", uuid).
		// Load Menus > Recipes > Ingredients > Product > NutritionFacts
		Relation("Menus.Recipes.Ingredients.Product.NutritionFacts").
		// Load Menus > Ingredients > Product > NutritionFacts
		Relation("Menus.Ingredients.Product.NutritionFacts").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("daily menu")
		}
		return nil, err
	}

	return model.DailyMenuToDomainDailyMenu(&dailyMenu), nil
}

func (x *DailyMenusRepository) GetDailyMenuByName(ctx context.Context, name string) (*domain.DailyMenu, error) {
	var menu model.DailyMenu

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

	return model.DailyMenuToDomainDailyMenu(&menu), nil
}
