package config

import (
	"context"
	"fmt"
	"gin-api/src/user/dtos/createUserDto"
	userRepositories "gin-api/src/user/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func Seed(client *mongo.Client) {

	userCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	userRepository := userRepositories.NewMongoUserRepository(userCollection)

	users, _ := userRepository.FindMany(context.Background(), 1, 10)
	if len(users) == 0 {
		firstUser := createUserDto.CreateUsersDto{
			Name:     "admin",
			Email:    "admin@fin.com",
			Password: "123456",
		}

		_, err := userRepository.Create(context.Background(), firstUser)
		if err != nil {
			return
		}

		fmt.Println("created the first database user")
	}
}
