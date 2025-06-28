package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	if price <= 0 {
		return nil, errors.New("price must be greather then 0")
	}
	if name == "" {
		return nil, errors.New("need name")
	}

	return &Product{
		ID:        uuid.NewString(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (p *Product) ChangePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}
	p.Price = price
	return nil
}

func (p *Product) ChangeName(name string) error {
	if name == "" {
		return errors.New("need name")
	}
	p.Name = name
	return nil
}
