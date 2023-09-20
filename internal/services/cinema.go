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
	storeCinemas, err := cs.CinemaRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}

	domainCinemas := make([]domain.Cinema, 0)
	for _, storeCinema := range storeCinemas {
		domainCinema, err := storeCinema.ToDomain()
		if err != nil {
			return nil, err
		}

		domainCinemas = append(domainCinemas, *domainCinema)
	}

	return domainCinemas, nil
}
