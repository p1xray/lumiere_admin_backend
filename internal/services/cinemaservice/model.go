package cinemaservice

import (
	"errors"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
)

// Входная модель кинотеатра
type CinemaInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address"`
}

// Возвращает доменную модель для создания
func (c *CinemaInput) ToCreateDomain() (*domain.Cinema, error) {
	if c == nil {
		return nil, errors.New("cinema input is nil")
	}
	return domain.NewCinemaToCreate(c.Name, c.Description, c.Address)
}
