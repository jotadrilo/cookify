package repositories

import (
	"context"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
)

type UnitsRepository interface {
	CreateUnit(context.Context, *domain.Unit) (*domain.Unit, error)
	ListUnits(context.Context) ([]*domain.Unit, error)
	GetUnitByID(context.Context, string) (*domain.Unit, error)
}

type UnitsRepositoryUnimpl struct{}

func (x *UnitsRepositoryUnimpl) CreateUnit(context.Context, *domain.Unit) (*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUnit")
}

func (x *UnitsRepositoryUnimpl) ListUnits(context.Context) ([]*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("ListUnits")
}

func (x *UnitsRepositoryUnimpl) GetUnitByID(context.Context, string) (*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("GetUnitByID")
}
