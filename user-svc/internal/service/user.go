package service

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/transport"
	"github.com/IvanMonichev/void-market-gin/user-svc/pkg/hash"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService interface {
	Create(ctx context.Context, dto transport.CreateUserDto) (*model.User, error)
	Find(ctx context.Context, id bson.ObjectID) (*model.User, error)
	Update(context.Context, *model.User, bson.ObjectID) (*model.User, error)
	Delete(ctx context.Context, id bson.ObjectID) error
}

type userService struct {
	repo repository.UserRepository
}

func New(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, dto transport.CreateUserDto) (*model.User, error) {
	hashedPassword, err := hash.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: hashedPassword,
	}

	return s.repo.Create(ctx, user)
}

func (s *userService) Find(ctx context.Context, id bson.ObjectID) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, user *model.User, id bson.ObjectID) (*model.User, error) {
	return s.repo.Update(ctx, user, id)
}

func (s *userService) Delete(ctx context.Context, id bson.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
