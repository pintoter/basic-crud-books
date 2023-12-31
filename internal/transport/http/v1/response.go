package v1

import (
	"github.com/pintoter/basic-crud-books/pkg/logger"

	"github.com/gin-gonic/gin"
)

func newResponse(c *gin.Context, statusCode int, msg any) {
	logger.Error(c.Request.RequestURI, "successfully")
	c.JSON(statusCode, msg)
}

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, err string) {
	logger.Error(c.Request.RequestURI, err)
	c.AbortWithStatusJSON(statusCode, errorResponse{Error: err})
}
