package service

import "api/test/catalog/internal/domain"

type ProductService interface {
	Save(name string, price float64) (string, error)
	Update(id string, name string, price float64) error
	FindById(id string) (*domain.Product, error)
	FindAll() []*domain.Product
	DeleteById(id string) error
}
