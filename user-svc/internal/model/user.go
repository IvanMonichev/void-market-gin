package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string        `json:"email"`
	Password  string        `json:"-"`
	Name      string        `json:"name"`
	UpdatedAt time.Time     `json:"updatedAt"`
	CreatedAt time.Time     `json:"createdAt"`
}
