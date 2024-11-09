package gin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/ports"
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

	if errors.Is(err, errorutils.ErrAlreadyExists) {
		m.Status = http.StatusConflict
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
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetProducts(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) PostProducts(c *gin.Context) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetProductsParamProductID(c *gin.Context, _ api.ParamProductID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserID(c *gin.Context, userID api.ParamUserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDDailyMenus(c *gin.Context, userID api.ParamUserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDDailyMenusParamDailyMenuID(c *gin.Context, userID api.ParamUserID, dailyMenuID api.ParamDailyMenuID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDMenus(c *gin.Context, userID api.ParamUserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDMenusParamMenuID(c *gin.Context, userID api.ParamUserID, menuID api.ParamMenuID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDRecipes(c *gin.Context, userID api.ParamUserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) PostUsersParamUserIDRecipes(c *gin.Context, userID api.ParamUserID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) GetUsersParamUserIDRecipesParamRecipeID(c *gin.Context, userID api.ParamUserID, recipeID api.ParamRecipeID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

func (u UnimplementedGinController) PatchUsersParamUserIDRecipesParamRecipeID(c *gin.Context, userID api.ParamUserID, recipeID api.ParamRecipeID) {
	RestErrorHandler(c, errorutils.NewErrNotImplemented("endpoint"))
}

var _ ports.GinController = (*UnimplementedGinController)(nil)

type Controller struct {
	UnimplementedGinController

	Users      ports.UsersUseCase
	Products   ports.ProductsUseCase
	Recipes    ports.RecipesUseCase
	Menus      ports.MenusUseCase
	DailyMenus ports.DailyMenusUseCase
}

func NewController(
	users ports.UsersUseCase,
	products ports.ProductsUseCase,
	recipes ports.RecipesUseCase,
	menus ports.MenusUseCase,
	dailyMenus ports.DailyMenusUseCase,
) *Controller {
	return &Controller{
		Users:      users,
		Products:   products,
		Recipes:    recipes,
		Menus:      menus,
		DailyMenus: dailyMenus,
	}
}

var _ ports.GinController = (*Controller)(nil)
