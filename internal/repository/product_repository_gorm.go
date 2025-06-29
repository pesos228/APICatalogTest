package repository

import (
	"api/test/catalog/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func (g *gormProductRepository) DeleteById(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&domain.Product{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return result.Error
}

func (g *gormProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	var product []*domain.Product
	result := g.db.WithContext(ctx).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (g *gormProductRepository) FindById(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product
	result := g.db.WithContext(ctx).First(&product, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, ErrProductNotFound
	}
	return &product, nil
}

func (g *gormProductRepository) Save(ctx context.Context, product *domain.Product) error {
	return g.db.WithContext(ctx).Save(product).Error
}

func NewGormProductRepository(db *gorm.DB) ProductRepository {
	return &gormProductRepository{db: db}
}
