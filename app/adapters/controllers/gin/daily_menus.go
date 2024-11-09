package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
)

func (x *Controller) GetUsersParamUserIDDailyMenus(c *gin.Context, _ api.ParamUserID) {
	var ctx = c.Request.Context()

	items, err := x.DailyMenus.ListDailyMenus(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainDailyMenusToAPIDailyMenus(items))
}

func (x *Controller) GetUsersParamUserIDDailyMenusParamDailyMenuID(c *gin.Context, _ api.ParamUserID, dailyMenuID api.ParamDailyMenuID) {
	var ctx = c.Request.Context()

	v, err := x.DailyMenus.GetDailyMenuByUUID(ctx, dailyMenuID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainDailyMenuToAPIDailyMenuDetailed(v))
}
