package user

import (
	"context"

	"fiber-app/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUsers(filter bson.M) ([]User, error) {
	cursor, err := database.DB.Collection("users").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

    var users []User
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func saveUser(user *CreateUserDTO) error {
	_, err := database.DB.Collection("users").InsertOne(context.Background(), user)
	return err
}
