package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jotadrilo/cookify/adapters/controllers"
	"github.com/jotadrilo/cookify/adapters/repositories/pg"
	"github.com/jotadrilo/cookify/adapters/services"
	_ "github.com/proullon/ramsql/driver"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"runtime"
)

func initProdPostgresDB() (*sql.DB, error) {
	db := otelsql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", cfg.PostgresHost, cfg.PostgresPort)),
		pgdriver.WithUser(cfg.PostgresUser),
		pgdriver.WithPassword(cfg.PostgresPass),
		pgdriver.WithDatabase(cfg.PostgresDatabase),
		pgdriver.WithInsecure(cfg.PostgresInsecure),
	))

	// Good practices in production
	// https://bun.uptrace.dev/guide/running-bun-in-production.html

	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxOpenConns)

	return db, nil
}

func initRamPostgresDB() (*sql.DB, error) {
	return sql.Open("ramsql", "test")
}

func initSQLDB() (*sql.DB, error) {
	if cfg.PostgresHost == "" {
		return initRamPostgresDB()
	}

	return initProdPostgresDB()
}

func initController(ctx context.Context) (*controllers.RestController, error) {
	sqldb, err := initSQLDB()
	if err != nil {
		return nil, err
	}

	db, err := pg.NewDB(ctx, sqldb)
	if err != nil {
		return nil, err
	}

	productsRepo := pg.NewProductsRepository(&pg.ProductsRepositoryOptions{
		DB: db,
	})

	unitsRepo := pg.NewUnitsRepository(&pg.UnitsRepositoryOptions{
		DB: db,
	})

	productsSvc := services.NewProductsService(&services.ProductsServiceOptions{
		Products: productsRepo,
	})

	unitsSvc := services.NewUnitsService(&services.UnitsServiceOptions{
		Units: unitsRepo,
	})

	if err := Seed(ctx, unitsSvc, productsSvc); err != nil {
		return nil, err
	}

	return controllers.NewRestController(&controllers.RestControllerOptions{
		Products: productsSvc,
		Units:    unitsSvc,
	}), nil
}
