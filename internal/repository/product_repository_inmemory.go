package repository

import (
	"api/test/catalog/internal/domain"
	"errors"
	"sync"
)

type inMemoryProductRepository struct {
	mutex    sync.Mutex
	products map[string]*domain.Product
}

func (r *inMemoryProductRepository) Save(product *domain.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.products[product.ID] = product
	return nil
}

func (r *inMemoryProductRepository) FindById(id string) (*domain.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	product, ok := r.products[id]

	if !ok {
		return nil, errors.New("Product with id " + id + " not found")
	}

	return product, nil
}

func (r *inMemoryProductRepository) FindAll() ([]*domain.Product, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var res []*domain.Product
	for _, val := range r.products {
		res = append(res, val)
	}

	return res, nil
}

func (r *inMemoryProductRepository) DeleteById(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.products[id]
	if !ok {
		return errors.New("product with id " + id + " not found")
	}

	delete(r.products, id)
	return nil
}

func NewProductRepository() ProductRepository {
	return &inMemoryProductRepository{products: make(map[string]*domain.Product)}
}
