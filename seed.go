package main

import (
	"context"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/ports/services"
	_ "github.com/proullon/ramsql/driver"
)

var units = []*domain.Unit{
	{
		ID:        "3db74a85-7a00-410f-8dd3-4daa738c51f5",
		Name:      "Gram",
		Unit10em3: "mg",
		Unit1:     "g",
		Unit10e3:  "kg",
	},
	{
		ID:        "b9bfa87c-2adb-4cf7-b5a5-c056fd4045e6",
		Name:      "Litre",
		Unit10em3: "ml",
		Unit1:     "l",
	},
}

var products = []*domain.Product{
	{
		ID:   "1",
		Name: "chicken",
		Units: []*domain.Unit{
			units[0],
		},
	},
}

func Seed(ctx context.Context, unitsSvc services.UnitsService, productsSvc services.ProductsService) error {
	for _, unit := range units {
		if _, err := unitsSvc.CreateUnit(ctx, unit); err != nil {
			return err
		}
	}

	for _, product := range products {
		if _, err := productsSvc.CreateProduct(ctx, product); err != nil {
			return err
		}
	}

	return nil
}
