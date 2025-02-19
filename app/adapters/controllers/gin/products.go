package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/oapi"
)

func (x *Controller) PostProducts(c *gin.Context) {
	var (
		ctx  = c.Request.Context()
		body api.PostProductsJSONRequestBody
	)

	if err := c.BindJSON(&body); err != nil {
		RestErrorHandler(c, err)
		return
	}

	logger.Infof("Received %+v", body)

	var product = ProductRequestToDomainProduct(&body)

	if err := product.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	item, err := x.Products.CreateProduct(ctx, product)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, &api.ProductID{Uuid: oapi.UUID(item.UUID)})
}

func (x *Controller) GetProducts(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Products.ListProducts(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainProductsToAPIProducts(items))
}

func (x *Controller) GetProductsParamProductID(c *gin.Context, productID api.ParamProductID) {
	var ctx = c.Request.Context()

	v, err := x.Products.GetProductByUUID(ctx, productID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainProductToAPIProduct(v))
}

func (x *Controller) GetProductsParamProductIDEquivalents(c *gin.Context, productID api.ParamProductID) {
	var ctx = c.Request.Context()

	items, err := x.Products.GetProductEquivalentsByUUID(ctx, productID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainEquivalentProductsToAPIEquivalentProducts(items))
}
