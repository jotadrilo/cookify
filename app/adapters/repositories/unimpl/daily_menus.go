package unimpl

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type DailyMenusRepository struct {
	ports.Repository
}

var _ ports.DailyMenusRepository = (*DailyMenusRepository)(nil)

func (x *DailyMenusRepository) ListDailyMenus(context.Context) ([]*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("ListDailyMenus")
}

func (x *DailyMenusRepository) GetDailyMenuByUUID(context.Context, string) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("GetDailyMenuByUUID")
}

func (x *DailyMenusRepository) CreateUserDailyMenu(context.Context, string, *domain.DailyMenu) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserDailyMenu")
}

func (x *DailyMenusRepository) ListUserDailyMenus(context.Context, string) ([]*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserDailyMenus")
}

func (x *DailyMenusRepository) GetUserDailyMenuByUUID(context.Context, string, string) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserDailyMenuByUUID")
}
