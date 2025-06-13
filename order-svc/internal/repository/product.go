package repository

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"github.com/google/uuid"
)

type Product interface {
	Create(ctx context.Context, product *Product) (*model.Product, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Product, error)
}
