package main

import (
	"context"

	"github.com/jotadrilo/cookify/client/cmd"
	"github.com/jotadrilo/cookify/internal/logger"
)

func main() {
	var ctx = context.Background()

	if err := cmd.NewRootCmd().ExecuteContext(ctx); err != nil {
		logger.Fatalf("Execution failed: %s", err.Error())
	}
}
