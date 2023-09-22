package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	idParam = "id"
)

func GetParamFromRoute(c *gin.Context, name string) (int64, error) {
	routeId := c.Param(name)
	if routeId == "" {
		return 0, ErrInvalidIdParam
	}

	id, err := strconv.ParseInt(routeId, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetIdFromRoute(c *gin.Context) (int64, error) {
	return GetParamFromRoute(c, idParam)
}
