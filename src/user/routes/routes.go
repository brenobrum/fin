package routes

import (
	objectIdParam "gin-api/domain/dtos/objectId"
	"gin-api/domain/dtos/pagination"
	"gin-api/src/auth/middleware"
	"gin-api/src/user/dtos/createUserDto"
	"gin-api/src/user/dtos/updateUserDto"
	userRepositories "gin-api/src/user/repositories"
	"gin-api/src/user/usecases/createUsers"
	"gin-api/src/user/usecases/findAllUsers"
	"gin-api/src/user/usecases/findUsers"
	"gin-api/src/user/usecases/updateUsers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func SetupUserRoutes(router *gin.Engine, client *mongo.Client) {
	userCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	userRepository := userRepositories.NewMongoUserRepository(userCollection)

	// set the JwtMiddleware for auth
	users := router.Group("/users", middleware.JwtAuthMiddleware())

	// TODO test
	// create a user
	users.POST("", func(c *gin.Context) {
		newUser, valid := createUserDto.Validate(c, userCollection)
		if !valid {
			return
		}

		createUsers.CreateUsers(c, userRepository, newUser)
	})

	// TODO test
	// update a user
	users.PUT("/:id", func(c *gin.Context) {
		userId, validated := objectIdParam.Validate(c)
		if !validated {
			return
		}

		newUser, validated := updateUserDto.Validate(c, userCollection)
		if !validated {
			return
		}

		updateUsers.UpdateUsers(c, userRepository, newUser, userId)
	})

	// TODO test
	// get one user
	users.GET("/:id", func(c *gin.Context) {
		userId, validated := objectIdParam.Validate(c)
		if !validated {
			return
		}

		findUsers.FindUsers(c, userRepository, userId)
	})

	// TODO test
	// get all users
	users.GET("", func(c *gin.Context) {
		pag := pagination.Validate(c)

		findAllUsers.FindAllUsers(c, userRepository, pag)
	})

}
