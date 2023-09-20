package services

import (
	"context"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
	"github.com/p1xray/lumiere_admin_backend/internal/repositories"
)

// Сервис кинотеатров
type CinemaService struct {
	CinemaRepo repositories.Cinemas
}

// Возвращает новый сервис кинотеатров
func NewCinemaService(c repositories.Cinemas) *CinemaService {
	return &CinemaService{
		CinemaRepo: c,
	}
}

// Возвращает список кинотеатров
func (cs *CinemaService) GetList(ctx context.Context) ([]domain.Cinema, error) {
	return nil, nil
}
