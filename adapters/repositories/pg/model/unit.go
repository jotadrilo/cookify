package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/jotadrilo/cookify/domain"
	"github.com/uptrace/bun"
)

type Unit struct {
	bun.BaseModel `bun:"table:units"`

	ID        int64  `bun:"id,pk,autoincrement"`
	UUID      string `bun:"uuid,notnull"`
	Name      string `bun:"name,notnull"`
	Unit10em3 string `bun:"unit_10em3"`
	Unit1     string `bun:"unit_1,notnull"`
	Unit10e3  string `bun:"unit_10e3"`
}

var _ bun.BeforeAppendModelHook = (*Unit)(nil)

func (x *Unit) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if x.UUID == "" {
		x.UUID = uuid.NewString()
	}

	return nil
}

func UnitToDomainUnit(x *Unit) *domain.Unit {
	return &domain.Unit{
		ID:        x.UUID,
		Name:      x.Name,
		Unit10em3: x.Unit10em3,
		Unit1:     x.Unit1,
		Unit10e3:  x.Unit10e3,
	}
}

func DomainUnitToUnit(x *domain.Unit) *Unit {
	return &Unit{
		UUID:      x.ID,
		Name:      x.Name,
		Unit10em3: x.Unit10em3,
		Unit1:     x.Unit1,
		Unit10e3:  x.Unit10e3,
	}
}
