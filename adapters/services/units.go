package services

import (
	"context"
	"fmt"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/ports/repositories"
	"github.com/jotadrilo/cookify/ports/services"
)

type UnitsService struct {
	services.UnitsServiceUnimpl

	Units repositories.UnitsRepository
}

type UnitsServiceOptions struct {
	Units repositories.UnitsRepository
}

func NewUnitsService(opts *UnitsServiceOptions) *UnitsService {
	return &UnitsService{
		Units: opts.Units,
	}
}

func (x *UnitsService) CreateUnit(ctx context.Context, u *domain.Unit) (*domain.Unit, error) {
	un, err := x.Units.CreateUnit(ctx, u)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created unit %q", un.ID)

	return un, nil
}

func (x *UnitsService) ListUnits(ctx context.Context) ([]*domain.Unit, error) {
	return x.Units.ListUnits(ctx)
}

func (x *UnitsService) GetUnitByID(ctx context.Context, id string) (*domain.Unit, error) {
	return x.Units.GetUnitByID(ctx, id)
}
