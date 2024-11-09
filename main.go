package main

import (
	"context"

	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/seed"
)

func main() {
	if err := mainE(context.Background()); err != nil {
		logger.Fatal(err)
	}
}

func mainE(ctx context.Context) error {
	ctl, err := initGinController(ctx)
	if err != nil {
		return err
	}

	if err := seed.Run(ctx, ctl); err != nil {
		return err
	}

	return serve(ctx, ctl)
}
