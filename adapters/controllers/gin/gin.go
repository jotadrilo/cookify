package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

func RestErrorHandler(c *gin.Context, err error) {
	var m = api.Error{
		Title:  err.Error(),
		Status: http.StatusInternalServerError,
	}

	if errors.Is(err, errorutils.ErrNotImplemented) {
		m.Status = http.StatusNotImplemented
	}

	if errors.Is(err, errorutils.ErrNotFound) {
		m.Status = http.StatusNotFound
	}

	if errors.Is(err, errorutils.ErrBadRequest) {
		m.Status = http.StatusBadRequest
	}

	c.JSON(m.Status, m)

	_ = c.AbortWithError(m.Status, err)
}

type UnimplementedGinController struct {
	ports.Controller
}

func (u UnimplementedGinController) GetAdminUsers(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetAdminUsers"))
}

func (u UnimplementedGinController) GetProducts(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetProducts"))
}

func (u UnimplementedGinController) GetProductsProductId(c *gin.Context, _ api.ProductID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetProductsProductId"))
}

func (u UnimplementedGinController) GetProductsProductIdNutritionFacts(c *gin.Context, _ api.ProductID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetProductsProductIdNutritionFacts"))
}

func (u UnimplementedGinController) GetUsersUserId(c *gin.Context, userID api.UserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserId"))
}

func (u UnimplementedGinController) GetUsersUserIdDailyMenus(c *gin.Context, userID api.UserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdDailyMenus"))
}

func (u UnimplementedGinController) GetUsersUserIdDailyMenusDailyMenuId(c *gin.Context, userID api.UserID, dailyMenuID api.DailyMenuID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdDailyMenusDailyMenuId"))
}

func (u UnimplementedGinController) GetUsersUserIdMenus(c *gin.Context, userID api.UserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdMenus"))
}

func (u UnimplementedGinController) GetUsersUserIdMenusMenuId(c *gin.Context, userID api.UserID, menuID api.MenuID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdMenusMenuId"))
}

func (u UnimplementedGinController) GetUsersUserIdRecipes(c *gin.Context, userID api.UserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdRecipes"))
}

func (u UnimplementedGinController) GetUsersUserIdRecipesRecipeId(c *gin.Context, userID api.UserID, recipeID api.RecipeID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("GetUsersUserIdRecipesRecipeId"))
}

var _ ports.GinController = (*UnimplementedGinController)(nil)

type Controller struct {
	UnimplementedGinController

	*UsersController
	*ProductsController
	*RecipesController
	*MenusController
	*DailyMenusController
}

func NewController(
	users *UsersController,
	products *ProductsController,
	recipes *RecipesController,
	menus *MenusController,
	dailyMenus *DailyMenusController,
) *Controller {
	return &Controller{
		UsersController:      users,
		ProductsController:   products,
		RecipesController:    recipes,
		MenusController:      menus,
		DailyMenusController: dailyMenus,
	}
}

var _ ports.GinController = (*Controller)(nil)

func (u Controller) GetAdminUsers(c *gin.Context) {
	u.UsersController.GetUsers(c)
}

func (u Controller) GetProducts(c *gin.Context) {
	u.ProductsController.GetProducts(c)
}

func (u Controller) GetProductsProductId(c *gin.Context, productID api.ProductID) {
	u.ProductsController.GetProductsProductId(c, productID)
}

func (u Controller) GetProductsProductIdNutritionFacts(c *gin.Context, productID api.ProductID) {
	u.ProductsController.GetProductsProductIdNutritionFacts(c, productID)
}

func (u Controller) GetUsersUserId(c *gin.Context, userID api.UserID) {
	u.UsersController.GetUsersUserId(c, userID)
}

func (u Controller) GetUsersUserIdDailyMenus(c *gin.Context, userID api.UserID) {
	u.DailyMenusController.GetUsersUserIdDailyMenus(c, userID)
}

func (u Controller) GetUsersUserIdDailyMenusDailyMenuId(c *gin.Context, userID api.UserID, dailyMenuID api.DailyMenuID) {
	u.DailyMenusController.GetUsersUserIdDailyMenusDailyMenuId(c, userID, dailyMenuID)
}

func (u Controller) GetUsersUserIdMenus(c *gin.Context, userID api.UserID) {
	u.MenusController.GetUsersUserIdMenus(c, userID)
}

func (u Controller) GetUsersUserIdMenusMenuId(c *gin.Context, userID api.UserID, menuID api.MenuID) {
	u.MenusController.GetUsersUserIdMenusMenuId(c, userID, menuID)
}

func (u Controller) GetUsersUserIdRecipes(c *gin.Context, userID api.UserID) {
	u.RecipesController.GetUsersUserIdRecipes(c, userID)
}

func (u Controller) GetUsersUserIdRecipesRecipeId(c *gin.Context, userID api.UserID, recipeID api.RecipeID) {
	u.RecipesController.GetUsersUserIdRecipesRecipeId(c, userID, recipeID)
}
