package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func BadRequest(c *gin.Context, err string) {
	c.AbortWithError(http.StatusBadRequest, errors.New(err))
}

func NotFound(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
}
