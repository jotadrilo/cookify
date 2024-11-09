package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jotadrilo/cookify/domain"
	"github.com/jotadrilo/cookify/errorutils"
	"net/http"
)

func RestErrorHandler(c *gin.Context, err error) {
	var m = domain.Error{
		Error: err.Error(),
		Code:  http.StatusInternalServerError,
	}

	if errors.Is(err, errorutils.ErrNotImplemented) {
		m.Code = http.StatusNotImplemented
	}

	if errors.Is(err, errorutils.ErrNotFound) {
		m.Code = http.StatusNotFound
	}

	if errors.Is(err, errorutils.ErrBadRequest) {
		m.Code = http.StatusBadRequest
	}

	c.JSON(m.Code, m)

	_ = c.AbortWithError(m.Code, err)
}
