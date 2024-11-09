package main

import (
	"context"
	"flag"

	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/server/config"
	"github.com/jotadrilo/cookify/server/seed"
)

func main() {
	config.BindFlags(flag.CommandLine)

	flag.Parse()

	if err := mainE(context.Background()); err != nil {
		logger.Fatal(err)
	}
}

func mainE(ctx context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}

	ctl, err := initGinController(ctx)
	if err != nil {
		return err
	}

	if err := seed.Run(ctx, ctl); err != nil {
		return err
	}

	return serve(ctx, ctl)
}
