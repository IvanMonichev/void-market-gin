package service

import (
	"context"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/repository"
	"github.com/IvanMonichev/void-market-gin/user-svc/pkg/hash"
)

type CreateUserInput struct {
	Email    string
	Name     string
	Password string
}

type UserService interface {
	Create(ctx context.Context, input CreateUserInput) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func New(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, input CreateUserInput) (*model.User, error) {
	hashedPassword, err := hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: hashedPassword,
	}

	return s.repo.Create(ctx, user)
}
