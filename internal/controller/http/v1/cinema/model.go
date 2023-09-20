package cinema

import (
	"time"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Модель кинотеатра для запросов
type Cinema struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Cinema) FillFrom(cinema domain.Cinema) {
	c.Id = cinema.Id()
	c.Name = cinema.Name()
	c.Description = cinema.Description()
	c.Address = cinema.Address()
	c.CreatedAt = cinema.CreatedAt()
	c.UpdatedAt = cinema.UpdatedAt()
}
