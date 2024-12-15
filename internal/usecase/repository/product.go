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
	query := `INSERT INTO products (name, description, price, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`

	var id int64
	result, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Price)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), err
}
func (r *productRepo) GetByID(ctx context.Context, id uint64) (*entity.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = ?`
	p := &entity.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}
func (r *productRepo) Update(ctx context.Context, p *entity.Product) error {
	query := `UPDATE products SET name=?, description=?, price=?, updated_at=NOW() WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.Price, p.ID)
	return err
}
func (r *productRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM products WHERE id = ?`
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
