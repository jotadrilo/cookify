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

type DailyMenusRepository struct {
	unimpl.DailyMenusRepository

	root string
}

var _ ports.DailyMenusRepository = (*DailyMenusRepository)(nil)

type DailyMenusRepositoryOptions struct {
	Root string
}

func NewDailyMenusRepository(opts *DailyMenusRepositoryOptions) *DailyMenusRepository {
	return &DailyMenusRepository{
		root: opts.Root,
	}
}

func (x *DailyMenusRepository) getFile() string {
	return filepath.Join(x.root, "daily_menus.json")
}

func (x *DailyMenusRepository) ListDailyMenus(_ context.Context) ([]*domain.DailyMenu, error) {
	items, err := decodeJSON[*model.DailyMenu](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map(items, model.DailyMenuToDomainDailyMenu), nil
}

func (x *DailyMenusRepository) GetDailyMenuByUUID(_ context.Context, uuid string) (*domain.DailyMenu, error) {
	items, err := decodeJSON[*model.DailyMenu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == uuid {
			return model.DailyMenuToDomainDailyMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("daily menu")
}

func (x *DailyMenusRepository) CreateUserDailyMenu(ctx context.Context, userID string, v *domain.DailyMenu) (*domain.DailyMenu, error) {
	if vv, err := x.GetUserDailyMenuByName(ctx, userID, v.Name); err == nil {
		return vv, nil
	}

	var (
		vv = model.DomainDailyMenuToDailyMenu(v)
	)

	vv.UUID = uuid.New().String()
	vv.UserUUID = userID

	if err := appendToJSON(x.getFile(), vv); err != nil {
		return nil, err
	}

	return model.DailyMenuToDomainDailyMenu(vv), nil
}

func (x *DailyMenusRepository) ListUserDailyMenus(_ context.Context, userID string) ([]*domain.DailyMenu, error) {
	items, err := decodeJSON[*model.DailyMenu](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map(slices.Select(items, model.User{UUID: userID}.OwnsDailyMenu), model.DailyMenuToDomainDailyMenu), nil
}

func (x *DailyMenusRepository) GetUserDailyMenuByUUID(_ context.Context, userID string, dailyMenuID string) (*domain.DailyMenu, error) {
	items, err := decodeJSON[*model.DailyMenu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == dailyMenuID && (model.User{UUID: userID}).OwnsDailyMenu(item) {
			return model.DailyMenuToDomainDailyMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("daily menu")
}

func (x *DailyMenusRepository) GetUserDailyMenuByName(_ context.Context, userID string, dailyMenuName string) (*domain.DailyMenu, error) {
	items, err := decodeJSON[*model.DailyMenu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Name == dailyMenuName && (model.User{UUID: userID}).OwnsDailyMenu(item) {
			return model.DailyMenuToDomainDailyMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("daily menu")
}
