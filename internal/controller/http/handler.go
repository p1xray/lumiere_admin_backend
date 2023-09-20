package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/config"
)

// Обработчик запросов http сервера
type Handler struct {
}

// Возвращает новый обработчик запросов http сервера
func NewHandler() *Handler {
	return &Handler{}
}

// Инициализирует обработчик запросов http сервера
func (h *Handler) Init(cfg config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
