package services

import (
	"context"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Сервис кинотеатров
type CinemaService struct {
}

// Возвращает новый сервис кинотеатров
func NewCinemaService() *CinemaService {
	return &CinemaService{}
}

// Возвращает список кинотеатров
func (cs *CinemaService) GetList(ctx context.Context) ([]domain.Cinema, error) {
	return nil, nil
}
