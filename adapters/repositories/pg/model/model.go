package model

import (
	"context"
	"embed"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Register(ctx context.Context, db *bun.DB) error {
	m := migrate.NewMigrations()
	if err := m.Discover(migrations); err != nil {
		return err
	}

	// Initialize migration tables
	migrator := migrate.NewMigrator(db, m)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	// Run migrations
	mg, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if !mg.IsZero() {
		fmt.Printf("Migrations applied successfully: %s\n", mg.String())
	}

	// Register M2M model so bun can better recognize m2m relation.
	// This should be done before you use the model for the first time.
	db.RegisterModel((*ProductToUnit)(nil))

	return nil
}
