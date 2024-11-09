package pg

import (
	"context"
	"database/sql"
	"github.com/jotadrilo/cookify/adapters/repositories/pg/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDB(ctx context.Context, sqldb *sql.DB) (*bun.DB, error) {
	// Good practices in production
	// https://bun.uptrace.dev/guide/running-bun-in-production.html
	db := bun.NewDB(sqldb, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := model.Register(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}
