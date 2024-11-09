package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
)

func (x *Controller) GetAdminUsers(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Users.ListUsers(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainUsersToAPIUsers(items))
}

func (x *Controller) GetUsersParamUserID(c *gin.Context, userID api.ParamUserID) {
	var ctx = c.Request.Context()

	v, err := x.Users.GetUserByUUID(ctx, userID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainUserToAPIUser(v))
}
