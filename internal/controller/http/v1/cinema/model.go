package cinema

import (
	"errors"
	"time"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Модель кинотеатра для запросов
type Cinema struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Заполняет данные модели кинотеатра для запросов данными доменной модели
func (c *Cinema) FillFrom(cinema domain.Cinema) error {
	if c == nil {
		return errors.New("cinema is nil")
	}

	c.Id = cinema.Id()
	c.Name = cinema.Name()
	c.Description = cinema.Description()
	c.Address = cinema.Address()
	c.CreatedAt = cinema.CreatedAt()
	c.UpdatedAt = cinema.UpdatedAt()

	return nil
}
