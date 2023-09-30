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

// @Summary Получить список кинотеатров
// @Tags Cinema
// @Description Получить список кинотеатров
// @ModuleID getCinemaList
// @Produce json
// @Success	200 {object}	server.dataResponse[CinemaList]
// @Failure	500	{object}	server.response
// @Router	/api/v1/cinema [get]
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

	response := CinemaList{
		Cinemas: responseCinemas,
	}

	server.SuccessResponse(c, response)
}

// @Summary Получить подробности кинотеатра
// @Tags Cinema
// @Description Получить подробности кинотеатра по идентификатору
// @ModuleID getCinemaDetails
// @Produce json
// @Param id path int64 true "Идентификатор кинотеатра"
// @Success	200 {object}	server.dataResponse[Cinema]
// @Failure	500	{object}	server.response
// @Router	/api/v1/cinema/{id} [get]
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

// @Summary Создать новый кинотеатр
// @Tags Cinema
// @Description Создать новый кинотеатр
// @ModuleID createCinema
// @Produce json
// @Param input body cinemaservice.CinemaInput true "Входные параметры для создания кинотеатра"
// @Success	200 {object}	server.dataResponse[bool]
// @Failure	500	{object}	server.response
// @Router	/api/v1/cinema [post]
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

// @Summary Редактировать кинотеатр
// @Tags Cinema
// @Description Редактировать кинотеатр по идентификатору
// @ModuleID updateCinema
// @Produce json
// @Param id path int64 true "Идентификатор кинотеатра"
// @Param input body cinemaservice.CinemaInput true "Входные параметры для редактирования кинотеатра"
// @Success	200 {object}	server.dataResponse[bool]
// @Failure	500	{object}	server.response
// @Router	/api/v1/cinema/{id} [put]
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

// @Summary Удалить кинотеатр
// @Tags Cinema
// @Description Удалить кинотеатр по идентификатору
// @ModuleID deleteCinema
// @Produce json
// @Param id path int64 true "Идентификатор кинотеатра"
// @Success	200 {object}	server.dataResponse[bool]
// @Failure	500	{object}	server.response
// @Router	/api/v1/cinema/{id} [delete]
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
