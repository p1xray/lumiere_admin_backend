package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/controller/http/v1/cinema"
)

// Обработчик запросов к API v1
type Handler struct {
}

// Возвращает новый обработчик запросов к API v1
func NewHandler() *Handler {
	return &Handler{}
}

// Инициализирует обработчик запросов к API v1
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		cinema.InitCinemaRoutes(v1)
	}
}
