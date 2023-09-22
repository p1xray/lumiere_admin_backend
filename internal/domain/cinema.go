package domain

import (
	"fmt"
	"time"
)

// Доменная модель кинотеатра
type Cinema struct {
	id          int64
	name        string
	description string
	address     string
	createdAt   time.Time
	updatedAt   time.Time
}

// Возвращает новую доменную модель кинотеатра
func NewCinema(id int64, name, description, address string, createdAt, updatedAt time.Time) (*Cinema, error) {
	if id < 1 {
		return nil, fmt.Errorf("%w: cinema id is required", ErrRequired)
	}

	if name == "" {
		return nil, fmt.Errorf("%w: cinema name is required", ErrRequired)
	}

	if address == "" {
		return nil, fmt.Errorf("%w: cinema address is required", ErrRequired)
	}

	cinema := &Cinema{
		id:          id,
		name:        name,
		description: description,
		address:     address,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}

	return cinema, nil
}

// Возвращает новую доменную модель кинотеатра для создания
func NewCinemaToCreate(name, description, address string) (*Cinema, error) {
	if name == "" {
		return nil, fmt.Errorf("%w: cinema name is required", ErrRequired)
	}

	if address == "" {
		return nil, fmt.Errorf("%w: cinema address is required", ErrRequired)
	}

	cinema := &Cinema{
		name:        name,
		description: description,
		address:     address,
	}

	return cinema, nil
}

// Возвращает идентификатор кинотеатра
func (c *Cinema) Id() int64 {
	return c.id
}

// Возвращает наименование кинотеатра
func (c *Cinema) Name() string {
	return c.name
}

// Возвращает описание кинотеатра
func (c *Cinema) Description() string {
	return c.description
}

// Возвращает адрес кинотеатра
func (c *Cinema) Address() string {
	return c.address
}

// Возвращает дату и время создания записи кинотеатра
func (c *Cinema) CreatedAt() time.Time {
	return c.createdAt
}

// Возвращает дату и время обновления записи кинотеатра
func (c *Cinema) UpdatedAt() time.Time {
	return c.updatedAt
}
