package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/jotadrilo/cookify/domain"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID    int64   `bun:"id,pk,autoincrement"`
	UUID  string  `bun:"uuid,notnull"`
	Name  string  `bun:"name,notnull"`
	Units []*Unit `bun:"m2m:product_to_units,join:Product=Unit"`
}

var _ bun.BeforeAppendModelHook = (*Product)(nil)

func (r *Product) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.UUID == "" {
		r.UUID = uuid.NewString()
	}

	return nil
}

func ProductToDomainProduct(x *Product) *domain.Product {
	var units []*domain.Unit

	for _, u := range x.Units {
		units = append(units, UnitToDomainUnit(u))
	}

	return &domain.Product{
		ID:    x.UUID,
		Name:  x.Name,
		Units: units,
	}
}

func DomainProductToProduct(x *domain.Product) *Product {
	var units []*Unit

	for _, u := range x.Units {
		units = append(units, DomainUnitToUnit(u))
	}

	return &Product{
		UUID:  x.ID,
		Name:  x.Name,
		Units: units,
	}
}
