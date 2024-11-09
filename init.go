package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"

	ginctl "github.com/jotadrilo/cookify/adapters/controllers/gin"
	fsrepo "github.com/jotadrilo/cookify/adapters/repositories/fs"
	pgrepo "github.com/jotadrilo/cookify/adapters/repositories/pg"
	"github.com/jotadrilo/cookify/adapters/usecases"
	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/config"
)

func initPostgresDB(ctx context.Context) (*bun.DB, error) {
	var cfg = config.Default()

	sqldb := otelsql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", cfg.PostgresHost, cfg.PostgresPort)),
		pgdriver.WithUser(cfg.PostgresUser),
		pgdriver.WithPassword(cfg.PostgresPass),
		pgdriver.WithDatabase(cfg.PostgresDatabase),
		pgdriver.WithInsecure(cfg.PostgresInsecure),
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

	return pgrepo.NewDB(ctx, sqldb)
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

	if cfg.FsRoot != "" {
		usersRepo = fsrepo.NewUsersRepository(&fsrepo.UsersRepositoryOptions{
			Root: cfg.FsRoot,
		})

		productsRepo = fsrepo.NewProductsRepository(&fsrepo.ProductsRepositoryOptions{
			Root: cfg.FsRoot,
		})

		recipesRepo = fsrepo.NewRecipesRepository(&fsrepo.RecipesRepositoryOptions{
			Root: cfg.FsRoot,
		})

		menusRepo = fsrepo.NewMenusRepository(&fsrepo.MenusRepositoryOptions{
			Root: cfg.FsRoot,
		})

		dailyMenusRepo = fsrepo.NewDailyMenusRepository(&fsrepo.DailyMenusRepositoryOptions{
			Root: cfg.FsRoot,
		})
	} else {
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
		ginctl.NewUsersController(usersUC),
		ginctl.NewProductsController(productsUC),
		ginctl.NewRecipesController(recipesUC),
		ginctl.NewMenusController(menusUC),
		ginctl.NewDailyMenusController(dailyMenusUC),
	), nil
}

func serve(_ context.Context, ctl *ginctl.Controller) error {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.Use(
		ginctl.ZapLogger(),
		ginctl.ZapRecovery(),
	)

	api.RegisterHandlersWithOptions(engine, ctl, api.GinServerOptions{
		BaseURL: "/api/v1",
	})

	return engine.Run(config.Default().ServerAddress)
}
