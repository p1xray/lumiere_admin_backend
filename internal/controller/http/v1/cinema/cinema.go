package cinema

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/server"
	"github.com/p1xray/lumiere_admin_backend/internal/services"
)

// Роуты для кинотеатров
type CinemaRoutes struct {
	CinemaService services.Cinemas
}

// Инициализация поутов для кинотеатров
func InitCinemaRoutes(api *gin.RouterGroup, s *services.Services) {
	cr := &CinemaRoutes{
		CinemaService: s.Cinemas,
	}

	cinema := api.Group("/cinema")
	{
		cinema.GET("", cr.getCinemaList)
		cinema.GET("/:id", cr.getCinemaDetails)
	}
}

func (cr *CinemaRoutes) getCinemaList(c *gin.Context) {
	cinemas, err := cr.CinemaService.GetList(c.Request.Context())
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	responseCinemas := make([]Cinema, 0)
	for _, cinema := range cinemas {
		responseCinema := Cinema{}
		responseCinema.FillFrom(cinema)
		responseCinemas = append(responseCinemas, responseCinema)
	}

	server.SuccessResponse(c, responseCinemas)
}

func (cr *CinemaRoutes) getCinemaDetails(c *gin.Context) {
	id, err := server.GetIdFromRoute(c)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	cinema, err := cr.CinemaService.GetDetails(c.Request.Context(), id)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	response := Cinema{}
	response.FillFrom(*cinema)

	server.SuccessResponse(c, response)
}
