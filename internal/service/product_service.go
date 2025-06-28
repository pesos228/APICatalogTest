package service

import (
	"api/test/catalog/internal/domain"
	"context"
)

type ProductService interface {
	Save(ctx context.Context, name string, price float64) (string, error)
	Update(ctx context.Context, id string, name string, price float64) error
	FindById(ctx context.Context, id string) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, error)
	DeleteById(ctx context.Context, id string) error
}
