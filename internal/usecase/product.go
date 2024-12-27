package usecase

import (
	"context"
	"github.com/dariuszdroba/go-from-template/internal/entity"
	"github.com/dariuszdroba/go-from-template/internal/usecase/repository"
)

type ProductUseCase interface {
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

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: r}
}

func (uc *productUseCase) Create(ctx context.Context, p *entity.Product) (uint64, error) {
	return uc.repo.Create(ctx, p)
}

func (uc *productUseCase) GetByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return uc.repo.GetByID(ctx, id)
}
func (uc *productUseCase) Update(ctx context.Context, p *entity.Product) error {
	return uc.repo.Update(ctx, p)
}
func (uc *productUseCase) Delete(ctx context.Context, id uint64) error {
	return uc.repo.Delete(ctx, id)
}
func (uc *productUseCase) List(ctx context.Context) ([]*entity.Product, error) {
	return uc.repo.List(ctx)
}
func (uc *productUseCase) GetProductHistory(ctx context.Context, id uint64) (*entity.Product, []*entity.ProductHistory, error) {
	return uc.repo.GetProductHistory(ctx, id)
}
func (uc *productUseCase) GetHighestPrice(ctx context.Context, id uint64) (*entity.ProductMaxValue, error) {
	return uc.repo.GetHighestPrice(ctx, id)
}
func (uc *productUseCase) GetTimeDiff(ctx context.Context, id uint64) ([]*entity.TimeDiff, error) {
	return uc.repo.GetTimeDiff(ctx, id)
}
func (uc *productUseCase) GetByDate(ctx context.Context, id uint64, rd *entity.ReferenceDate) (*entity.ProductHistory, error) {
	return uc.repo.GetByDate(ctx, id, rd)
}
