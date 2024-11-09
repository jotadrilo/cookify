package pg

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/jotadrilo/cookify/adapters/repositories/pg/model"
	"github.com/jotadrilo/cookify/internal/config"
)

func NewDB(ctx context.Context, sqldb *sql.DB) (*bun.DB, error) {
	// Good practices in production
	// https://bun.uptrace.dev/guide/running-bun-in-production.html
	db := bun.NewDB(sqldb, pgdialect.New(),
		bun.WithDiscardUnknownColumns(),
	)

	var cfg = config.Default()

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(cfg.BunVerbose),
	))

	if err := model.Register(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}
