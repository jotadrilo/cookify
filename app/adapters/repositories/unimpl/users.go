package unimpl

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type UsersRepository struct {
	ports.Repository
}

var _ ports.UsersRepository = (*UsersRepository)(nil)

func (x *UsersRepository) CreateUser(context.Context, *domain.User) (*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUser")
}

func (x *UsersRepository) ListUsers(context.Context) ([]*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("ListUsers")
}

func (x *UsersRepository) GetUserByUUID(context.Context, string) (*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserByUUID")
}
