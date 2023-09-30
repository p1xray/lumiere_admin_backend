package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p1xray/lumiere_admin_backend/internal/config"
	v1 "github.com/p1xray/lumiere_admin_backend/internal/controller/http/v1"
	"github.com/p1xray/lumiere_admin_backend/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/p1xray/lumiere_admin_backend/docs"
)

// Обработчик запросов http сервера
type Handler struct {
	Services *services.Services
}

// Возвращает новый обработчик запросов http сервера
func NewHandler(s *services.Services) *Handler {
	return &Handler{
		Services: s,
	}
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
	initSwagger(router)

	return router
}

// Инициализирует обработчик запросов к API
func (h *Handler) initAPI(router *gin.Engine) {
	v1 := v1.NewHandler(h.Services)
	api := router.Group("/api")
	{
		v1.Init(api)
	}
}

// Инициализирует документацию swagger
func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
