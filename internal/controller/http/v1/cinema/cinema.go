package cinema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Роуты для кинотеатров
type CinemaRoutes struct {
}

// Инициализация поутов для кинотеатров
func InitCinemaRoutes(api *gin.RouterGroup) {
	cr := &CinemaRoutes{}

	cinema := api.Group("/cinema")
	{
		cinema.GET("", cr.getCinemaList)
	}
}

func (cr *CinemaRoutes) getCinemaList(c *gin.Context) {
	e := &Cinema{
		Id:          1,
		Name:        "cinema list",
		Description: "cinema list",
		Address:     "cinema list",
	}
	c.JSON(http.StatusOK, e)
}
