package repository

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
	FindById(ctx context.Context, orderId uint) (*model.Order, error)
	Update(ctx context.Context, order *model.Order, id uint) (*model.Order, error)
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context, offset int64, limit int64) ([]model.Order, int64, error)
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	if err := r.db.WithContext(ctx).Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *GormOrderRepository) FindById(ctx context.Context, orderId uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.WithContext(ctx).First(&order, orderId).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *GormOrderRepository) Update(ctx context.Context, order *model.Order, id uint) (*model.Order, error) {
	order.ID = id
	if err := r.db.WithContext(ctx).Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *GormOrderRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormOrderRepository) GetAll(ctx context.Context, offset int64, limit int64) ([]model.Order, int64, error) {
	var (
		orders []model.Order
		total  int64
	)

	if err := r.db.WithContext(ctx).Model(&model.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Применяем пагинацию
	if err := r.db.WithContext(ctx).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}
