package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/controller/http/v1/cinema"
	"github.com/p1xray/lumiere_admin_backend/internal/services"
)

// Обработчик запросов к API v1
type Handler struct {
	Services *services.Services
}

// Возвращает новый обработчик запросов к API v1
func NewHandler(s *services.Services) *Handler {
	return &Handler{
		Services: s,
	}
}

// Инициализирует обработчик запросов к API v1
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		cinema.InitCinemaRoutes(v1, h.Services)
	}
}
