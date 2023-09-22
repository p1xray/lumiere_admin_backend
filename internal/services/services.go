package services

import (
	"context"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
	"github.com/p1xray/lumiere_admin_backend/internal/repositories"
	"github.com/p1xray/lumiere_admin_backend/internal/services/cinemaservice"
)

// Интерфейс сервиса кинотеатров
type Cinemas interface {
	// Возвращает список кинотеатров
	GetList(ctx context.Context) ([]domain.Cinema, error)

	// Возвращает подробности кинотеатра
	GetDetails(ctx context.Context, id int64) (*domain.Cinema, error)

	// Создает новый кинотеатр
	Create(ctx context.Context, inp *cinemaservice.CinemaInput) error

	// Обновляет данные о кинотеатре
	Update(ctx context.Context, id int64, inp *cinemaservice.CinemaInput) error
}

// Сервисы
type Services struct {
	Cinemas Cinemas
}

// Зависимости сервисов
type Deps struct {
	Repos *repositories.Repositories
}

// Возвращает сервисы с зависимостями
func NewServices(deps Deps) *Services {
	return &Services{
		Cinemas: cinemaservice.NewService(deps.Repos.Cinemas),
	}
}
