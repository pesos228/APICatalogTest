package repository

import (
	"api/test/catalog/internal/domain"
	"context"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) error
	FindById(ctx context.Context, id string) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, error)
	DeleteById(ctx context.Context, id string) error
}
