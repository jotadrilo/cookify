package usecases

import (
	"context"

	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
)

type UnimplementedUsersUseCase struct {
	ports.UseCase
}

var _ ports.UsersUseCase = (*UnimplementedUsersUseCase)(nil)

func (x *UnimplementedUsersUseCase) CreateUser(context.Context, *domain.User) (*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUser")
}

func (x *UnimplementedUsersUseCase) ListUsers(context.Context) ([]*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("ListUsers")
}

func (x *UnimplementedUsersUseCase) GetUserByUUID(context.Context, string) (*domain.User, error) {
	return nil, errorutils.NewErrNotImplemented("GetUserByUUID")
}

type UsersUseCase struct {
	UnimplementedUsersUseCase

	Users ports.UsersRepository
}

type UsersUseCaseOptions struct {
	Users ports.UsersRepository
}

func NewUsersUseCase(opts *UsersUseCaseOptions) *UsersUseCase {
	return &UsersUseCase{
		Users: opts.Users,
	}
}

func (x *UsersUseCase) CreateUser(ctx context.Context, v *domain.User) (*domain.User, error) {
	vv, err := x.Users.CreateUser(ctx, v)
	if err != nil {
		logger.Errorf("Cannot create user %v: %s", v, err.Error())
		return nil, errorutils.NewErrNotCreated("user")
	}

	logger.Infof("Created user %q", vv.UUID)

	return vv, nil
}

func (x *UsersUseCase) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return x.Users.ListUsers(ctx)
}

func (x *UsersUseCase) GetUserByUUID(ctx context.Context, uuid string) (*domain.User, error) {
	return x.Users.GetUserByUUID(ctx, uuid)
}
