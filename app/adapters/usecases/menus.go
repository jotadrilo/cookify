package usecases

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
)

type UnimplementedMenusUseCase struct {
	ports.UseCase
}

var _ ports.MenusUseCase = (*UnimplementedMenusUseCase)(nil)

func (x *UnimplementedMenusUseCase) ListMenus(context.Context) ([]*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("ListMenus")
}

func (x *UnimplementedMenusUseCase) GetMenuByUUID(context.Context, string) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("GetMenuByUUID")
}

func (x *UnimplementedMenusUseCase) CreateUserMenu(context.Context, string, *domain.Menu) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserMenu")
}

func (x *UnimplementedMenusUseCase) ListUserMenus(context.Context, string) ([]*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserMenus")
}

func (x *UnimplementedMenusUseCase) GetUserMenuByUUID(context.Context, string, string) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserMenuByUUID")
}

type MenusUseCase struct {
	UnimplementedMenusUseCase

	Menus ports.MenusRepository
}

type MenusUseCaseOptions struct {
	Menus ports.MenusRepository
}

func NewMenusUseCase(opts *MenusUseCaseOptions) *MenusUseCase {
	return &MenusUseCase{
		Menus: opts.Menus,
	}
}

func (x *MenusUseCase) ListMenus(ctx context.Context) ([]*domain.Menu, error) {
	return x.Menus.ListMenus(ctx)
}

func (x *MenusUseCase) GetMenuByUUID(ctx context.Context, menuID string) (*domain.Menu, error) {
	return x.Menus.GetMenuByUUID(ctx, menuID)
}

func (x *MenusUseCase) CreateUserMenu(ctx context.Context, userID string, v *domain.Menu) (*domain.Menu, error) {
	vv, err := x.Menus.CreateUserMenu(ctx, userID, v)
	if err != nil {
		logger.Errorf("Cannot create menu %v: %s", v, err.Error())
		return nil, errorutils.NewErrNotCreated("menu")
	}

	logger.Infof("Created menu %q", vv.UUID)

	return vv, nil
}

func (x *MenusUseCase) ListUserMenus(ctx context.Context, userID string) ([]*domain.Menu, error) {
	return x.Menus.ListUserMenus(ctx, userID)
}

func (x *MenusUseCase) GetUserMenuByUUID(ctx context.Context, userID string, menuID string) (*domain.Menu, error) {
	return x.Menus.GetUserMenuByUUID(ctx, userID, menuID)
}
