package services

import (
	"context"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Интерфейс сервиса кинотеатров
type Cinemas interface {
	GetList(ctx context.Context) ([]domain.Cinema, error)
}

// Сервисы
type Services struct {
	Cinemas Cinemas
}

// Зависимости сервисов
type Deps struct {
}

// Возвращает сервисы с зависимостями
func NewServices(deps Deps) *Services {
	return &Services{
		Cinemas: NewCinemaService(),
	}
}
