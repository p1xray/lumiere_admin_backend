package server

import (
	"context"
	"net/http"

	"github.com/p1xray/lumiere_admin_backend/internal/config"
)

// Сервер
type Server struct {
	httpServer *http.Server
}

// Возвращает новый сервер
func NewServer(cfg config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.HttpAddr,
			Handler: handler,
		},
	}
}

// Запускает сервер
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Останавливает сервер
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
