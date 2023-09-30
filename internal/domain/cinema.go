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

	cinema := &Cinema{
		id:          id,
		name:        name,
		description: description,
		address:     address,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}

	if err := cinema.validateCinema(); err != nil {
		return nil, err
	}

	return cinema, nil
}

// Возвращает новую доменную модель кинотеатра для создания
func NewCinemaToCreate(name, description, address string) (*Cinema, error) {
	cinema := &Cinema{
		name:        name,
		description: description,
		address:     address,
	}

	if err := cinema.validateCinema(); err != nil {
		return nil, err
	}

	return cinema, nil
}

// Обновляет поля доменной модели
func (c *Cinema) UpdateCinema(name, description, address string) error {
	c.name = name
	c.description = description
	c.address = address
	c.updatedAt = time.Now()

	if err := c.validateCinema(); err != nil {
		return err
	}

	return nil
}

// Валидация полей доменной модели
func (c *Cinema) validateCinema() error {
	if c.name == "" {
		return fmt.Errorf("%w: cinema name is required", ErrRequired)
	}

	if c.address == "" {
		return fmt.Errorf("%w: cinema address is required", ErrRequired)
	}

	return nil
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
