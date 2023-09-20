package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/config"
	v1 "github.com/p1xray/lumiere_admin_backend/internal/controller/http/v1"
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

	h.initAPI(router)

	return router
}

// Инициализирует обработчик запросов к API
func (h *Handler) initAPI(router *gin.Engine) {
	v1 := v1.NewHandler()
	api := router.Group("/api")
	{
		v1.Init(api)
	}
}
