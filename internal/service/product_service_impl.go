package service

import (
	"api/test/catalog/internal/domain"
	"api/test/catalog/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func (p *productService) Update(id string, name string, price float64) error {
	product, err := p.repo.FindById(id)
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

	err = p.repo.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) DeleteById(id string) error {
	err := p.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) FindAll() []*domain.Product {
	products, err := p.repo.FindAll()
	if err != nil {
		return nil
	}
	return products
}

func (p *productService) FindById(id string) (*domain.Product, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productService) Save(name string, price float64) (string, error) {
	product, err := domain.NewProduct(name, price)
	if err != nil {
		return "", err
	}
	err = p.repo.Save(product)
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
