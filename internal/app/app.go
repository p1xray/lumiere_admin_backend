package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/p1xray/lumiere_admin_backend/internal/config"
	controller "github.com/p1xray/lumiere_admin_backend/internal/controller/http"
	"github.com/p1xray/lumiere_admin_backend/internal/repositories"
	"github.com/p1xray/lumiere_admin_backend/internal/server"
	"github.com/p1xray/lumiere_admin_backend/internal/services"
	"github.com/p1xray/lumiere_admin_backend/pkg/postgres"
)

// Запускает приложение
func Run(cfg config.Config) {
	// БД
	pg, err := postgres.New(cfg.DatabaseUrl, postgres.MaxPoolSize(cfg.PoolMax))
	if err != nil {
		log.Fatalf("Postgres New Error: %v", err)
	}
	defer pg.Close()

	// Репозитории
	repos := repositories.NewRepositories(repositories.Deps{
		Postgres: pg,
	})

	// Сервисы
	services := services.NewServices(services.Deps{
		Repos: repos,
	})

	// Обработчики запросов
	handlers := controller.NewHandler(services)

	// Сервер
	srv := server.NewServer(cfg, handlers.Init(cfg))

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Stop(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HttpAddr)

	// Запуск http сервера
	if err := srv.Run(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped
}
