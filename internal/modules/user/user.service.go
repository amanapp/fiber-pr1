package user

import "go.mongodb.org/mongo-driver/bson"

func ListUsers(filter bson.M) ([]User, error) {
	return FindUsers(filter)
}

func CreateUser(user *CreateUserDTO) error {
	return saveUser(user)
}
