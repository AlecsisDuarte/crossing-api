package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

///Ok responds with a 200 Status Code with the object as the body of the response with a
///JSON structure
func Ok(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

///BadRequest responds with a 401 Status Code with the object as the body of the response with a
///JSON structure
func BadRequest(c *gin.Context, err string) {
	c.AbortWithError(http.StatusBadRequest, errors.New(err))
}

///BadRequest responds with a 404 Status Code with the object as the body of the response with a
///JSON structure
func NotFound(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
}

///BadRequest responds with a 501 Status Code with the object as the body of the response with a
///JSON structure
func InternalError(c *gin.Context, err string) {
	c.AbortWithError(http.StatusInternalServerError, errors.New(err))
}
