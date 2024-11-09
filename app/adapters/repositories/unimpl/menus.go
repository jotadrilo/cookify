package unimpl

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type MenusRepository struct {
	ports.Repository
}

var _ ports.MenusRepository = (*MenusRepository)(nil)

func (x *MenusRepository) ListMenus(context.Context) ([]*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("ListMenus")
}

func (x *MenusRepository) GetMenuByUUID(context.Context, string) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("GetMenuByUUID")
}

func (x *MenusRepository) CreateUserMenu(context.Context, string, *domain.Menu) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUserMenu")
}

func (x *MenusRepository) ListUserMenus(context.Context, string) ([]*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("ListUserMenus")
}

func (x *MenusRepository) GetUserMenuByUUID(context.Context, string, string) (*domain.Menu, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserMenuByUUID")
}
