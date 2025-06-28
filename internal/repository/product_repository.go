package repository

import "api/test/catalog/internal/domain"

type ProductRepository interface {
	Save(product *domain.Product) error
	FindById(id string) (*domain.Product, error)
	FindAll() ([]*domain.Product, error)
	DeleteById(id string) error
}
