package domain

import "fmt"

// Доменная модель кинотеатра
type Cinema struct {
	id          int64
	name        string
	description string
	address     string
}

// Возвращает новую доменную модель кинотеатра
func NewCinema(id int64, name, description, address string) (*Cinema, error) {
	if id < 1 {
		return nil, fmt.Errorf("%w: cinema id is required", ErrRequired)
	}

	if name == "" {
		return nil, fmt.Errorf("%w: cinema name is required", ErrRequired)
	}

	cinema := &Cinema{
		id:          id,
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
