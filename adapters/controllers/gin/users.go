package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type UsersController struct {
	Users ports.UsersUseCase
}

func NewUsersController(users ports.UsersUseCase) *UsersController {
	return &UsersController{
		Users: users,
	}
}

func (x *UsersController) CreateUser(c *gin.Context) {
	var (
		ctx  = c.Request.Context()
		user domain.User
	)

	if err := c.BindJSON(&user); err != nil {
		RestErrorHandler(c, err)
		return
	}

	if err := user.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Users.CreateUser(ctx, &user)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *UsersController) GetUsers(c *gin.Context) {
	var ctx = c.Request.Context()

	items, err := x.Users.ListUsers(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainUsersToAPIUsers(items))
}

func (x *UsersController) GetUsersUserId(c *gin.Context, userID api.UserID) {
	var ctx = c.Request.Context()

	v, err := x.Users.GetUserByUUID(ctx, userID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainUserToAPIUser(v))
}
