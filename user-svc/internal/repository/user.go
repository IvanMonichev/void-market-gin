package repository

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	FindByID(ctx context.Context, id bson.ObjectID) (*model.User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) UserRepository {
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = bson.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id bson.ObjectID) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
