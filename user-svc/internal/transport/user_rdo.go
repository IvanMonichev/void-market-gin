package transport

import "github.com/IvanMonichev/void-market-gin/user-svc/internal/model"

type UserRdo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUserRdo(u *model.User) UserRdo {
	return UserRdo{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name:  u.Name,
	}
}
