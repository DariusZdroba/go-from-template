package repository

import (
	"context"
	"database/sql"
	"github.com/dariuszdroba/go-from-template/internal/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, p *entity.Product) (uint64, error)
	GetByID(ctx context.Context, id uint64) (*entity.Product, error)
	Update(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context) ([]*entity.Product, error)
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(ctx context.Context, p *entity.Product) (uint64, error) {
	query := `INSERT INTO products (name, description, price, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	var id uint64
	err := r.db.QueryRowContext(ctx, query, p.Name, p.Description, p.Price).Scan(&id)
	return id, err
}
func (r *productRepo) GetByID(ctx context.Context, id uint64) (*entity.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	p := &entity.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}
func (r *productRepo) Update(ctx context.Context, p *entity.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, updated_at=NOW() WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Price, p.ID)
	return err
}
func (r *productRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func (r *productRepo) List(ctx context.Context) ([]*entity.Product, error) {
	query := `SELECT * FROM products`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		p := &entity.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
