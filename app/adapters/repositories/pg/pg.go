package pg

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/jotadrilo/cookify/app/adapters/repositories/pg/model"
)

func NewDB(ctx context.Context, sqldb *sql.DB, verbose bool) (*bun.DB, error) {
	// Good practices in production
	// https://bun.uptrace.dev/guide/running-bun-in-production.html
	db := bun.NewDB(sqldb, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(verbose),
	))

	if err := model.Register(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}
