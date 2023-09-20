package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message,omitempty"`
} // @name Response

type dataResponse[T any] struct {
	IsSuccess bool `json:"is_success"`
	Data      T    `json:"data,omitempty"`
} // @name DataResponse

func SuccessResponse[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, dataResponse[T]{IsSuccess: true, Data: data})
}

func ErrorResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusOK, response{IsSuccess: false, Message: message})
}
