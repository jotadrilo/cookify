package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type MenusController struct {
	Menus ports.MenusUseCase
}

func NewMenusController(recipes ports.MenusUseCase) *MenusController {
	return &MenusController{
		Menus: recipes,
	}
}

func (x *MenusController) CreateMenu(c *gin.Context, userID api.UserID) {
	var (
		ctx  = c.Request.Context()
		menu domain.Menu
	)

	if err := c.BindJSON(&menu); err != nil {
		RestErrorHandler(c, err)
		return
	}

	if err := menu.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Menus.CreateUserMenu(ctx, userID.String(), &menu)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *MenusController) GetUsersUserIdMenus(c *gin.Context, _ api.UserID) {
	var ctx = c.Request.Context()

	items, err := x.Menus.ListMenus(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainMenusToAPIMenus(items))
}

func (x *MenusController) GetUsersUserIdMenusMenuId(c *gin.Context, _ api.UserID, menuID api.MenuID) {
	var ctx = c.Request.Context()

	v, err := x.Menus.GetMenuByUUID(ctx, menuID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainMenuToAPIMenuDetailed(v))
}
