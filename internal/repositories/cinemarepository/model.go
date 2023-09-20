package cinemarepository

import (
	"errors"
	"time"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Модель кинотеатра для репозитория
type Cinema struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Возвращает доменную модель порта, созданную на основе модели порта для репозитория
func (c *Cinema) ToDomain() (*domain.Cinema, error) {
	if c == nil {
		return nil, errors.New("store cinema is nil")
	}
	return domain.NewCinema(c.Id, c.Name, c.Description, c.Address, c.CreatedAt, c.UpdatedAt)
}
