package cinema

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/server"
	"github.com/p1xray/lumiere_admin_backend/internal/services"
	"github.com/p1xray/lumiere_admin_backend/internal/services/cinemaservice"
)

// Роуты для кинотеатров
type CinemaRoutes struct {
	CinemaService services.Cinemas
}

// Инициализация поутов для кинотеатров
func InitRoutes(api *gin.RouterGroup, s *services.Services) {
	cr := &CinemaRoutes{
		CinemaService: s.Cinemas,
	}

	cinema := api.Group("/cinema")
	{
		cinema.GET("", cr.getList)
		cinema.GET("/:id", cr.getDetails)
		cinema.POST("", cr.create)
		cinema.PUT("/:id", cr.update)
		cinema.DELETE("/:id", cr.delete)
	}
}

func (cr *CinemaRoutes) getList(c *gin.Context) {
	cinemas, err := cr.CinemaService.GetList(c.Request.Context())
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	responseCinemas := make([]Cinema, 0)
	for _, cinema := range cinemas {
		responseCinema := Cinema{}
		if err := responseCinema.FillFrom(cinema); err != nil {
			server.ErrorResponse(c, err.Error())
			return
		}
		responseCinemas = append(responseCinemas, responseCinema)
	}

	server.SuccessResponse(c, responseCinemas)
}

func (cr *CinemaRoutes) getDetails(c *gin.Context) {
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
	if err := response.FillFrom(*cinema); err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	server.SuccessResponse(c, response)
}

func (cr *CinemaRoutes) create(c *gin.Context) {
	inp, err := server.GetInputFromBody[cinemaservice.CinemaInput](c)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	if err := cr.CinemaService.Create(c.Request.Context(), inp); err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	server.SuccessResponse(c, true)
}

func (cr *CinemaRoutes) update(c *gin.Context) {
	id, err := server.GetIdFromRoute(c)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	inp, err := server.GetInputFromBody[cinemaservice.CinemaInput](c)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	if err := cr.CinemaService.Update(c.Request.Context(), id, inp); err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	server.SuccessResponse(c, true)
}

func (cr *CinemaRoutes) delete(c *gin.Context) {
	id, err := server.GetIdFromRoute(c)
	if err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	if err = cr.CinemaService.Delete(c.Request.Context(), id); err != nil {
		server.ErrorResponse(c, err.Error())
		return
	}

	server.SuccessResponse(c, true)
}
