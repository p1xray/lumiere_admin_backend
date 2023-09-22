package cinemaservice

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
func NewService(c repositories.Cinemas) *CinemaService {
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

// Возвращает подробности кинотеатра
func (cs *CinemaService) GetDetails(ctx context.Context, id int64) (*domain.Cinema, error) {
	storeCinema, err := cs.CinemaRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	domainCinema, err := storeCinema.ToDomain()
	if err != nil {
		return nil, err
	}

	return domainCinema, nil
}

// Создает новый кинотеатр
func (cs *CinemaService) Create(ctx context.Context, inp *CinemaInput) error {
	domainCinema, err := inp.ToCreateDomain()
	if err != nil {
		return err
	}

	if err = cs.CinemaRepo.Create(ctx, domainCinema); err != nil {
		return err
	}

	return nil
}

// Обновляет данные о кинотеатре
func (cs *CinemaService) Update(ctx context.Context, id int64, inp *CinemaInput) error {
	storeCinema, err := cs.CinemaRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	domainCinema, err := storeCinema.ToDomain()
	if err != nil {
		return err
	}

	if err = domainCinema.UpdateCinema(inp.Name, inp.Description, inp.Address); err != nil {
		return err
	}

	if err = cs.CinemaRepo.Update(ctx, domainCinema); err != nil {
		return err
	}

	return nil
}

// Удаляет данные о кинотеатре
func (cs *CinemaService) Delete(ctx context.Context, id int64) error {
	storeCinema, err := cs.CinemaRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if err = cs.CinemaRepo.Delete(ctx, storeCinema.Id); err != nil {
		return err
	}

	return nil
}
