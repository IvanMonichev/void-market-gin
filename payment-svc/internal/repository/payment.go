package repository

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/payment-svc/internal/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *model.Payment) error
	FindByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
}

type MongoPaymentRepository struct {
	collection *mongo.Collection
}

func NewMongoPaymentRepository(db *mongo.Database) *MongoPaymentRepository {
	return &MongoPaymentRepository{
		collection: db.Collection("payments"),
	}
}

func (r *MongoPaymentRepository) Save(ctx context.Context, payment *model.Payment) error {
	payment.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, payment)
	return err
}

func (r *MongoPaymentRepository) FindByOrderID(ctx context.Context, orderID string) (*model.Payment, error) {
	var result model.Payment
	err := r.collection.FindOne(ctx, bson.M{"orderId": orderID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
