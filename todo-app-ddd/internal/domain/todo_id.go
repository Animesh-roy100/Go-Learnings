package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoID struct {
	value primitive.ObjectID
}

func NewTodoID() TodoID {
	return TodoID{value: primitive.NewObjectID()}
}

func (id TodoID) String() string {
	return id.value.Hex()
}

func TodoIDFromString(s string) (TodoID, error) {
	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return TodoID{}, errors.New("invalid TodoID")
	}
	return TodoID{value: objectID}, nil
}
