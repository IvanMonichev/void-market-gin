package transport

import (
	"github.com/IvanMonichev/void-market-gin/user-svc/internal/model"
	"time"
)

type UserRdo struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserRdo(u *model.User) UserRdo {
	return UserRdo{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
