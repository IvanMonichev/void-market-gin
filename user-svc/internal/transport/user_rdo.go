package transport

import "github.com/IvanMonichev/void-market-gin/user-svc/internal/model"

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUserRdo(u *model.User) UserResponse {
	return UserResponse{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name:  u.Name,
	}
}
