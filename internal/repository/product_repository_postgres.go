package repository

import (
	"api/test/catalog/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresProductRepository struct {
	db *pgxpool.Pool
}

var ErrProductNotFound = errors.New("product not found")

func (p *postgresProductRepository) DeleteById(ctx context.Context, id string) error {
	query := `DELETE FROM products WHERE id = $1`
	commandTag, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (p *postgresProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	query := `SELECT id, name, price, created_at FROM products`
	row, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(row, pgx.RowToAddrOfStructByName[domain.Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *postgresProductRepository) FindById(ctx context.Context, id string) (*domain.Product, error) {
	query := `SELECT id, name, price, created_at FROM products WHERE id = $1`
	row, err := p.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	product, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Product])

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}

		return nil, err
	}

	return &product, nil
}

func (p *postgresProductRepository) Save(ctx context.Context, product *domain.Product) error {
	query := `
        INSERT INTO products (id, name, price, created_at)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            price = EXCLUDED.price
    `
	_, err := p.db.Exec(ctx, query, product.ID, product.Name, product.Price, product.CreatedAt)
	return err
}

func NewPostgresProductRepository(db *pgxpool.Pool) ProductRepository {
	return &postgresProductRepository{db: db}
}
