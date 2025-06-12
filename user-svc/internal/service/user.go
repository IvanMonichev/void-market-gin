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
