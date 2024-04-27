package models

import (
	"fmt"
	"gin-api/domain/utils/copyType"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type ReadUser struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

func ResponseUser(user User) (readUser ReadUser) {

	err := copyType.CopyType(&user, &readUser)
	if err != nil {
		fmt.Println(err)

		return ReadUser{}
	}

	return readUser
}

func ResponseUsers(users []User) (readUsers []ReadUser) {
	for _, user := range users {
		readUser := ReadUser{}
		err := copyType.CopyType(&user, &readUser)
		if err != nil {
			fmt.Println(err)
		}

		readUsers = append(readUsers, readUser)
	}

	return readUsers
}
