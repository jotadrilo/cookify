package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type ProductsController struct {
	Products ports.ProductsUseCase
}

func NewProductsController(products ports.ProductsUseCase) *ProductsController {
	return &ProductsController{
		Products: products,
	}
}

func (x *ProductsController) CreateProduct(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		product domain.Product
	)

	if err := c.BindJSON(&product); err != nil {
		RestErrorHandler(c, err)
		return
	}

	if err := product.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Products.CreateProduct(ctx, &product)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *ProductsController) GetProducts(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Products.ListProducts(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainProductsToAPIProducts(items))
}

func (x *ProductsController) GetProductsProductId(c *gin.Context, productID api.ProductID) {
	var ctx = c.Request.Context()

	v, err := x.Products.GetProductByUUID(ctx, productID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainProductToAPIProductDetailed(v))
}

func (x *ProductsController) GetProductsProductIdNutritionFacts(c *gin.Context, productID api.ProductID) {
	var ctx = c.Request.Context()

	v, err := x.Products.GetProductNutritionFactsByUUID(ctx, productID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainNutritionFactsToAPINutritionFacts(v))
}
