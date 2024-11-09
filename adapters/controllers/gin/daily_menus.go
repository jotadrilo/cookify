package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type DailyMenusController struct {
	DailyMenus ports.DailyMenusUseCase
}

func NewDailyMenusController(recipes ports.DailyMenusUseCase) *DailyMenusController {
	return &DailyMenusController{
		DailyMenus: recipes,
	}
}

func (x *DailyMenusController) CreateUserDailyMenu(c *gin.Context, userID api.UserID) {
	var (
		ctx  = c.Request.Context()
		menu domain.DailyMenu
	)

	if err := c.BindJSON(&menu); err != nil {
		RestErrorHandler(c, err)
		return
	}

	if err := menu.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.DailyMenus.CreateUserDailyMenu(ctx, userID.String(), &menu)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *DailyMenusController) GetUsersUserIdDailyMenus(c *gin.Context, _ api.UserID) {
	var ctx = c.Request.Context()

	items, err := x.DailyMenus.ListDailyMenus(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainDailyMenusToAPIDailyMenus(items))
}

func (x *DailyMenusController) GetUsersUserIdDailyMenusDailyMenuId(c *gin.Context, _ api.UserID, dailyMenuID api.DailyMenuID) {
	var ctx = c.Request.Context()

	v, err := x.DailyMenus.GetDailyMenuByUUID(ctx, dailyMenuID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainDailyMenuToAPIDailyMenuDetailed(v))
}
