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

type UsersRepository struct {
	unimpl.UsersRepository

	root string
}

var _ ports.UsersRepository = (*UsersRepository)(nil)

type UsersRepositoryOptions struct {
	Root string
}

func NewUsersRepository(opts *UsersRepositoryOptions) *UsersRepository {
	return &UsersRepository{
		root: opts.Root,
	}
}

func (x *UsersRepository) getFile() string {
	return filepath.Join(x.root, "users.json")
}

func (x *UsersRepository) CreateUser(ctx context.Context, v *domain.User) (*domain.User, error) {
	if vv, err := x.GetUserByName(ctx, v.Name); err == nil {
		return vv, nil
	}

	var (
		vv = model.DomainUserToUser(v)
	)

	vv.UUID = uuid.New().String()

	if err := appendToJSON(x.getFile(), vv); err != nil {
		return nil, err
	}

	return model.UserToDomainUser(vv), nil
}

func (x *UsersRepository) ListUsers(_ context.Context) ([]*domain.User, error) {
	items, err := decodeJSON[*model.User](x.getFile())
	if err != nil {
		return nil, err
	}
	return slices.Map[*model.User, *domain.User](items, model.UserToDomainUser), nil
}

func (x *UsersRepository) GetUserByUUID(_ context.Context, uuid string) (*domain.User, error) {
	items, err := decodeJSON[*model.User](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.UUID == uuid {
			return model.UserToDomainUser(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("user")
}

func (x *UsersRepository) GetUserByName(_ context.Context, name string) (*domain.User, error) {
	items, err := decodeJSON[*model.User](x.getFile())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Name == name {
			return model.UserToDomainUser(item), nil
		}
	}

	return nil, errorutils.NewErrNotFound("user")
}
