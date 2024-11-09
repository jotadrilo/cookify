package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/proullon/ramsql/driver"
	"log"
	"net/http"
)

func main() {
	if err := mainE(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func mainE(ctx context.Context) error {
	ctl, err := initController(ctx)
	if err != nil {
		return err
	}

	e := gin.New()

	e.Handle(http.MethodGet, "/products", ctl.ListProducts)
	e.Handle(http.MethodPost, "/products", ctl.CreateProduct)
	e.Handle(http.MethodGet, "/products/:id", ctl.GetProductByID)

	e.Handle(http.MethodGet, "/units", ctl.ListUnits)
	e.Handle(http.MethodPost, "/units", ctl.CreateUnit)
	e.Handle(http.MethodGet, "/units/:id", ctl.GetUnitByID)

	return e.Run()
}
