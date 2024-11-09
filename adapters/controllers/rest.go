package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
	"github.com/jotadrilo/cookify/ports/controllers"
	"github.com/jotadrilo/cookify/ports/services"
	"net/http"
)

type RestController struct {
	controllers.RestControllerUnimpl

	Products services.ProductsService
	Units    services.UnitsService
}

type RestControllerOptions struct {
	Products services.ProductsService
	Units    services.UnitsService
}

func NewRestController(opts *RestControllerOptions) *RestController {
	return &RestController{
		Products: opts.Products,
		Units:    opts.Units,
	}
}

func (x *RestController) CreateProduct(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		product domain.Product
	)

	if err := c.BindJSON(&product); err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	if err := product.Validate(); err != nil {
		controllers.RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Products.CreateProduct(ctx, &product)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *RestController) ListProducts(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Products.ListProducts(ctx)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (x *RestController) GetProductByID(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		id  = c.Param("id")
	)

	if id == "" {
		controllers.RestErrorHandler(c, errorutils.NewErrBadRequest("id is required"))
	}

	p, err := x.Products.GetProductByID(ctx, id)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, p)
}

func (x *RestController) CreateUnit(c *gin.Context) {
	var (
		ctx  = c.Request.Context()
		unit = domain.Unit{}
	)

	if err := c.BindJSON(&unit); err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	if err := unit.Validate(); err != nil {
		controllers.RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Units.CreateUnit(ctx, &unit)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *RestController) ListUnits(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Units.ListUnits(ctx)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (x *RestController) GetUnitByID(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		id  = c.Param("id")
	)

	if id == "" {
		controllers.RestErrorHandler(c, errorutils.NewErrBadRequest("id is required"))
	}

	p, err := x.Units.GetUnitByID(ctx, id)
	if err != nil {
		controllers.RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, p)
}
