package pg

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jotadrilo/cookify/adapters/repositories/pg/model"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
	"github.com/jotadrilo/cookify/ports/repositories"
	"github.com/uptrace/bun"
)

type UnitsRepository struct {
	repositories.UnitsRepositoryUnimpl

	db *bun.DB
}

type UnitsRepositoryOptions struct {
	DB *bun.DB
}

func NewUnitsRepository(opts *UnitsRepositoryOptions) *UnitsRepository {
	return &UnitsRepository{
		db: opts.DB,
	}
}

func (x *UnitsRepository) CreateUnit(ctx context.Context, u *domain.Unit) (*domain.Unit, error) {
	u.ID = uuid.NewString()

	var (
		q = x.db.NewInsert().Model(model.DomainUnitToUnit(u))
	)

	if _, err := q.Exec(ctx); err != nil {
		return nil, err
	}

	return u, nil
}

func (x *UnitsRepository) ListUnits(ctx context.Context) ([]*domain.Unit, error) {
	var (
		items []*domain.Unit
		units []*model.Unit
		q     = x.db.NewSelect().Model(&units)
	)

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	for _, u := range units {
		items = append(items, model.UnitToDomainUnit(u))
	}

	return items, nil
}

func (x *UnitsRepository) GetUnitByID(ctx context.Context, id string) (*domain.Unit, error) {
	var (
		unit model.Unit
		q    = x.db.NewSelect().Model(&unit).Where("uuid = ?", id)
	)

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorutils.NewErrNotFound("unit")
		}
		return nil, err
	}

	return model.UnitToDomainUnit(&unit), nil
}
