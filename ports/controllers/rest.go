package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jotadrilo/cookify/errorutils"
)

type RestController interface {
	CreateProduct(c *gin.Context)
	ListProducts(c *gin.Context)
	GetProductByID(c *gin.Context)

	CreateUnit(c *gin.Context)
	ListUnits(c *gin.Context)
	GetUnitByID(c *gin.Context)
}

type RestControllerUnimpl struct{}

func (x *RestControllerUnimpl) CreateProduct(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("CreateProduct"))
}

func (x *RestControllerUnimpl) ListProducts(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("ListProducts"))
}

func (x *RestControllerUnimpl) GetProductByID(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetProductByID"))
}

func (x *RestControllerUnimpl) CreateUnit(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("CreateUnit"))
}

func (x *RestControllerUnimpl) ListUnits(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("ListUnits"))
}

func (x *RestControllerUnimpl) GetUnitByID(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUnitByID"))
}
