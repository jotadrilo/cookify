package services

import (
	"context"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
)

type UnitsService interface {
	CreateUnit(context.Context, *domain.Unit) (*domain.Unit, error)
	ListUnits(context.Context) ([]*domain.Unit, error)
	GetUnitByID(context.Context, string) (*domain.Unit, error)
}

type UnitsServiceUnimpl struct{}

func (x *UnitsServiceUnimpl) CreateUnit(context.Context, *domain.Unit) (*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("CreateUnit")
}

func (x *UnitsServiceUnimpl) ListUnits(context.Context) ([]*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("ListUnits")
}

func (x *UnitsServiceUnimpl) GetUnitByID(context.Context, string) (*domain.Unit, error) {
	return nil, errorutils.NewErrNotImplemented("GetUnitByID")
}
