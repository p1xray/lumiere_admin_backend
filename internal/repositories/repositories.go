package repositories

import (
	"context"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
	"github.com/p1xray/lumiere_admin_backend/internal/repositories/cinemarepository"
	"github.com/p1xray/lumiere_admin_backend/pkg/postgres"
)

// Интерфейс репозитория кинотеатра
type Cinemas interface {
	// Возвращает список кинотеатров из хранилища
	GetList(ctx context.Context) ([]cinemarepository.Cinema, error)

	// Возвращает подробности кинотеатра
	GetDetails(ctx context.Context, id int64) (*cinemarepository.Cinema, error)

	// Создает новую запись кинотеатра в хранилище
	Create(ctx context.Context, cinema *domain.Cinema) error
}

// Репозитории
type Repositories struct {
	Cinemas Cinemas
}

// Зависимости репозиториев
type Deps struct {
	Postgres *postgres.Postgres
}

// Возвращает репозитории с зависимостями
func NewRepositories(deps Deps) *Repositories {
	return &Repositories{
		Cinemas: cinemarepository.NewCinemaRepository(deps.Postgres),
	}
}
