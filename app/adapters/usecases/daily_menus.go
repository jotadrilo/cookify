package usecases

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
)

type UnimplementedDailyMenusUseCase struct {
	ports.UseCase
}

var _ ports.DailyMenusUseCase = (*UnimplementedDailyMenusUseCase)(nil)

func (x *UnimplementedDailyMenusUseCase) ListDailyMenus(context.Context) ([]*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("ListDailyMenus")
}

func (x *UnimplementedDailyMenusUseCase) GetDailyMenuByUUID(context.Context, string) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("GetDailyMenuByUUID")
}

func (x *UnimplementedDailyMenusUseCase) CreateUserDailyMenu(context.Context, string, *domain.DailyMenu) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserDailyMenu")
}

func (x *UnimplementedDailyMenusUseCase) ListUserDailyMenus(context.Context, string) ([]*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserDailyMenus")
}

func (x *UnimplementedDailyMenusUseCase) GetUserDailyMenuByUUID(context.Context, string, string) (*domain.DailyMenu, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserDailyMenuByUUID")
}

type DailyMenusUseCase struct {
	UnimplementedDailyMenusUseCase

	DailyMenus ports.DailyMenusRepository
}

type DailyMenusUseCaseOptions struct {
	DailyMenus ports.DailyMenusRepository
}

func NewDailyMenusUseCase(opts *DailyMenusUseCaseOptions) *DailyMenusUseCase {
	return &DailyMenusUseCase{
		DailyMenus: opts.DailyMenus,
	}
}

func (x *DailyMenusUseCase) ListDailyMenus(ctx context.Context) ([]*domain.DailyMenu, error) {
	return x.DailyMenus.ListDailyMenus(ctx)
}

func (x *DailyMenusUseCase) GetDailyMenuByUUID(ctx context.Context, dailyMenuID string) (*domain.DailyMenu, error) {
	return x.DailyMenus.GetDailyMenuByUUID(ctx, dailyMenuID)
}

func (x *DailyMenusUseCase) CreateUserDailyMenu(ctx context.Context, userID string, v *domain.DailyMenu) (*domain.DailyMenu, error) {
	vv, err := x.DailyMenus.CreateUserDailyMenu(ctx, userID, v)
	if err != nil {
		logger.Errorf("Cannot create daily menu %v: %s", v, err.Error())
		return nil, errorutils.NewErrNotCreated("daily menu")
	}

	logger.Infof("Created daily menu %q", vv.UUID)

	return vv, nil
}

func (x *DailyMenusUseCase) ListUserDailyMenus(ctx context.Context, userID string) ([]*domain.DailyMenu, error) {
	return x.DailyMenus.ListUserDailyMenus(ctx, userID)
}

func (x *DailyMenusUseCase) GetUserDailyMenuByUUID(ctx context.Context, userID string, dailyMenuID string) (*domain.DailyMenu, error) {
	return x.DailyMenus.GetUserDailyMenuByUUID(ctx, userID, dailyMenuID)
}
