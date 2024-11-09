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

type MenusRepository struct {
	unimpl.MenusRepository

	root string
}

var _ ports.MenusRepository = (*MenusRepository)(nil)

type MenusRepositoryOptions struct {
	Root string
}

func NewMenusRepository(opts *MenusRepositoryOptions) *MenusRepository {
	return &MenusRepository{
		root: opts.Root,
	}
}

func (x *MenusRepository) getFile() string {
	return filepath.Join(x.root, "menus.json")
}

func (x *MenusRepository) ListMenus(_ context.Context) ([]*domain.Menu, error) {
	items, err := decodeJSON[*model.Menu](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map(items, model.MenuToDomainMenu), nil
}

func (x *MenusRepository) GetMenuByUUID(_ context.Context, menuID string) (*domain.Menu, error) {
	items, err := decodeJSON[*model.Menu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == menuID {
			return model.MenuToDomainMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("menu")
}

func (x *MenusRepository) CreateUserMenu(ctx context.Context, userID string, v *domain.Menu) (*domain.Menu, error) {
	if vv, err := x.GetMenuByName(ctx, userID, v.Name); err == nil {
		return vv, nil
	}

	var (
		vv = model.DomainMenuToMenu(v)
	)

	vv.UUID = uuid.New().String()
	vv.UserUUID = userID

	if err := appendToJSON(x.getFile(), vv); err != nil {
		return nil, err
	}

	return model.MenuToDomainMenu(vv), nil
}

func (x *MenusRepository) ListUserMenus(_ context.Context, userID string) ([]*domain.Menu, error) {
	items, err := decodeJSON[*model.Menu](x.getFile())
	if err != nil {
		return nil, err
	}

	return slices.Map(slices.Select(items, model.User{UUID: userID}.OwnsMenu), model.MenuToDomainMenu), nil
}

func (x *MenusRepository) GetUserMenuByUUID(_ context.Context, userID string, menuID string) (*domain.Menu, error) {
	items, err := decodeJSON[*model.Menu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == menuID && (model.User{UUID: userID}).OwnsMenu(item) {
			return model.MenuToDomainMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("menu")
}

func (x *MenusRepository) GetMenuByName(_ context.Context, userID string, menuName string) (*domain.Menu, error) {
	items, err := decodeJSON[*model.Menu](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Name == menuName && (model.User{UUID: userID}).OwnsMenu(item) {
			return model.MenuToDomainMenu(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("menu")
}
