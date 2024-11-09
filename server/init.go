package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"

	ginctl "github.com/jotadrilo/cookify/app/adapters/controllers/gin"
	fsrepo "github.com/jotadrilo/cookify/app/adapters/repositories/fs"
	pgrepo "github.com/jotadrilo/cookify/app/adapters/repositories/pg"
	"github.com/jotadrilo/cookify/app/adapters/usecases"
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/ports"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/server/config"
)

func initPostgresDB(ctx context.Context) (*bun.DB, error) {
	var cfg = config.Default()

	sqldb := otelsql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", cfg.Database.Postgres.Host, cfg.Database.Postgres.Port)),
		pgdriver.WithUser(cfg.Database.Postgres.User),
		pgdriver.WithPassword(cfg.Database.Postgres.Pass),
		pgdriver.WithDatabase(cfg.Database.Postgres.Database),
		pgdriver.WithInsecure(cfg.Database.Postgres.Insecure),
	))

	go func() {
		select {
		case <-ctx.Done():
			_ = sqldb.Close()
		}
	}()

	// Good practices in production
	// https://bun.uptrace.dev/guide/running-bun-in-production.html

	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)

	return pgrepo.NewDB(ctx, sqldb, cfg.Database.Postgres.BunVerbose)
}

func initGinController(ctx context.Context) (*ginctl.Controller, error) {
	var (
		cfg = config.Default()

		usersRepo      ports.UsersRepository
		productsRepo   ports.ProductsRepository
		recipesRepo    ports.RecipesRepository
		menusRepo      ports.MenusRepository
		dailyMenusRepo ports.DailyMenusRepository
	)

	if cfg.Database.FileSystem != nil {
		logger.Info("Configuring filesystem database...")

		usersRepo = fsrepo.NewUsersRepository(&fsrepo.UsersRepositoryOptions{
			Root: cfg.Database.FileSystem.Root,
		})

		productsRepo = fsrepo.NewProductsRepository(&fsrepo.ProductsRepositoryOptions{
			Root: cfg.Database.FileSystem.Root,
		})

		recipesRepo = fsrepo.NewRecipesRepository(&fsrepo.RecipesRepositoryOptions{
			Root: cfg.Database.FileSystem.Root,
		})

		menusRepo = fsrepo.NewMenusRepository(&fsrepo.MenusRepositoryOptions{
			Root: cfg.Database.FileSystem.Root,
		})

		dailyMenusRepo = fsrepo.NewDailyMenusRepository(&fsrepo.DailyMenusRepositoryOptions{
			Root: cfg.Database.FileSystem.Root,
		})
	} else {
		logger.Info("Configuring PostgreSQL database...")

		db, err := initPostgresDB(ctx)
		if err != nil {
			return nil, err
		}

		productsRepo = pgrepo.NewProductsRepository(&pgrepo.ProductsRepositoryOptions{
			DB: db,
		})

		recipesRepo = pgrepo.NewRecipesRepository(&pgrepo.RecipesRepositoryOptions{
			DB: db,
		})

		menusRepo = pgrepo.NewMenusRepository(&pgrepo.MenusRepositoryOptions{
			DB: db,
		})

		dailyMenusRepo = pgrepo.NewDailyMenusRepository(&pgrepo.DailyMenusRepositoryOptions{
			DB: db,
		})
	}

	usersUC := usecases.NewUsersUseCase(&usecases.UsersUseCaseOptions{
		Users: usersRepo,
	})

	productsUC := usecases.NewProductsUseCase(&usecases.ProductsUseCaseOptions{
		Products: productsRepo,
	})

	recipesUC := usecases.NewRecipesUseCase(&usecases.RecipesUseCaseOptions{
		Recipes: recipesRepo,
	})

	menusUC := usecases.NewMenusUseCase(&usecases.MenusUseCaseOptions{
		Menus: menusRepo,
	})

	dailyMenusUC := usecases.NewDailyMenusUseCase(&usecases.DailyMenusUseCaseOptions{
		DailyMenus: dailyMenusRepo,
	})

	return ginctl.NewController(
		usersUC,
		productsUC,
		recipesUC,
		menusUC,
		dailyMenusUC,
	), nil
}

func serve(_ context.Context, ctl *ginctl.Controller) error {
	var cfg = config.Default()

	var mode = gin.TestMode

	if cfg.Server.Gin != nil {
		mode = cfg.Server.Gin.Mode
	}

	gin.SetMode(mode)

	engine := gin.New()

	if err := engine.SetTrustedProxies(nil); err != nil {
		return err
	}

	engine.Use(
		ginctl.ZapLogger(),
		ginctl.ZapRecovery(),
	)

	api.RegisterHandlersWithOptions(engine, ctl, api.GinServerOptions{
		BaseURL: "/api/v1",
	})

	logger.Infof("Listening and serving GIN web API on %s", cfg.Server.Address)

	return engine.Run(cfg.Server.Address)
}
