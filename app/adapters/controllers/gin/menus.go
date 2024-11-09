package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
)

func (x *Controller) GetUsersParamUserIDMenus(c *gin.Context, _ api.ParamUserID) {
	var ctx = c.Request.Context()

	items, err := x.Menus.ListMenus(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainMenusToAPIMenus(items))
}

func (x *Controller) GetUsersParamUserIDMenusParamMenuID(c *gin.Context, _ api.ParamUserID, menuID api.ParamMenuID) {
	var ctx = c.Request.Context()

	v, err := x.Menus.GetMenuByUUID(ctx, menuID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainMenuToAPIMenuDetailed(v))
}
