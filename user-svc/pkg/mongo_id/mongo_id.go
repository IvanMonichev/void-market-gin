package mongo_id

import (
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var ErrInvalidObjectID = errors.New("invalid MongoDB ObjectID")

func Parse(idStr string) (bson.ObjectID, error) {
	if idStr == "" {
		return bson.NilObjectID, ErrInvalidObjectID
	}

	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return bson.NilObjectID, ErrInvalidObjectID
	}

	return objID, nil
}
