package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	Name      string        `bson:"name" json:"name"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}
