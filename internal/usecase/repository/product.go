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
	GetProductHistory(ctx context.Context, id uint64) (*entity.Product, []*entity.ProductHistory, error)
	GetHighestPrice(ctx context.Context, id uint64) (*entity.ProductMaxValue, error)
	GetTimeDiff(ctx context.Context, id uint64) ([]*entity.TimeDiff, error)
	GetByDate(ctx context.Context, id uint64, rd *entity.ReferenceDate) (*entity.ProductHistory, error)
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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	// Initial Product state to history
	queryHistoryInsert := `INSERT into product_history (product_id, name, description, price, created_at, valid_from, valid_to) SELECT id, name, description, price, created_at, NOW(), NOW() + INTERVAL 7 DAY FROM products WHERE id = ?`

	_, err = tx.ExecContext(ctx, queryHistoryInsert, p.ID)
	if err != nil {
		return err
	}

	// Update Product
	queryUpdateProduct := ` UPDATE products SET name = ?, description = ?, price = ?, updated_at = NOW() WHERE id = ?`
	_, err = tx.ExecContext(ctx, queryUpdateProduct, p.Name, p.Description, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
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

func (r *productRepo) GetProductHistory(ctx context.Context, id uint64) (*entity.Product, []*entity.ProductHistory, error) {
	productQuery := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = ?`
	historyQuery := `SELECT name, description, price, valid_from, valid_to FROM product_history WHERE product_id = ?`

	p := &entity.Product{}
	err := r.db.QueryRowContext(ctx, productQuery, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, nil, err
	}

	rows, err := r.db.QueryContext(ctx, historyQuery, id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var history []*entity.ProductHistory
	for rows.Next() {
		h := &entity.ProductHistory{}
		err := rows.Scan(&h.Name, &h.Description, &h.Price, &h.ValidFrom, &h.ValidTo)
		if err != nil {
			return nil, nil, err
		}
		history = append(history, h)
	}
	return p, history, nil
}

func (r *productRepo) GetHighestPrice(ctx context.Context, id uint64) (*entity.ProductMaxValue, error) {
	maxPriceQuery := `SELECT price, TIMEDIFF(valid_to, valid_from) AS duration FROM test.product_history WHERE product_id = ? ORDER BY duration DESC, price DESC LIMIT 1; `
	pMax := &entity.ProductMaxValue{}
	err := r.db.QueryRowContext(ctx, maxPriceQuery, id).Scan(&pMax.Price, &pMax.Duration)
	if err != nil {
		return nil, err
	}
	return pMax, nil
}
func (r *productRepo) GetTimeDiff(ctx context.Context, id uint64) ([]*entity.TimeDiff, error) {
	timeDiffQuery := `SELECT valid_from, valid_to, price FROM test.product_history WHERE product_id = ? ORDER BY valid_from, valid_to;`
	var tDiffs []*entity.TimeDiff
	rows, err := r.db.QueryContext(ctx, timeDiffQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		t := &entity.TimeDiff{}
		err := rows.Scan(&t.ValidFrom, &t.ValidTo, &t.Price)
		if err != nil {
			return nil, err
		}
		tDiffs = append(tDiffs, t)
	}
	return tDiffs, nil
}
func (r *productRepo) GetByDate(ctx context.Context, id uint64, rd *entity.ReferenceDate) (*entity.ProductHistory, error) {
	byDateQuery := `SELECT * FROM test.product_history WHERE product_id = ? AND ? BETWEEN valid_from AND valid_to ORDER BY valid_from DESC LIMIT 1`
	ph := &entity.ProductHistory{}
	err := r.db.QueryRowContext(ctx, byDateQuery, id, rd.DateTime).
		Scan(&ph.ID, &ph.ProductID, &ph.Name, &ph.Description, &ph.Price, &ph.ValidFrom, &ph.ValidTo)
	if err != nil {
		return nil, err
	}
	return ph, nil
}
