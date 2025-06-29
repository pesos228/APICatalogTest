package service

import (
	"api/test/catalog/internal/domain"
	"api/test/catalog/internal/repository"
	"context"
	"errors"
	"fmt"
)

type productService struct {
	repo repository.ProductRepository
}

func (p *productService) Update(ctx context.Context, id string, name string, price float64) error {
	product, err := p.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = product.ChangeName(name)
	if err != nil {
		return err
	}

	err = product.ChangePrice(price)
	if err != nil {
		return err
	}

	err = p.repo.Save(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) DeleteById(ctx context.Context, id string) error {
	err := p.repo.DeleteById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return fmt.Errorf("%w: product with id '%s'", repository.ErrProductNotFound, id)
		}
		return err
	}
	return nil
}

func (p *productService) FindAll(ctx context.Context) ([]*domain.Product, error) {
	products, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productService) FindById(ctx context.Context, id string) (*domain.Product, error) {
	product, err := p.repo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, fmt.Errorf("%w: product with id '%s'", repository.ErrProductNotFound, id)
		}
		return nil, err
	}
	return product, nil
}

func (p *productService) Save(ctx context.Context, name string, price float64) (string, error) {
	product, err := domain.NewProduct(name, price)
	if err != nil {
		return "", err
	}
	err = p.repo.Save(ctx, product)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}
