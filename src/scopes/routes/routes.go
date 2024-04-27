package routes

import (
	"gin-api/src/auth/middleware"
	createScopesDto "gin-api/src/scopes/dtos/createUserDto"
	scopeRepositories "gin-api/src/scopes/repositories"
	createScope "gin-api/src/scopes/usecases/createScopes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func SetupScopeRoutes(router *gin.Engine, client *mongo.Client) {
	scopeCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("scopes")
	scopeRepository := scopeRepositories.NewMongoScopeRepository(scopeCollection)

	// set the JwtMiddleware for auth
	scopes := router.Group("/scopes", middleware.JwtAuthMiddleware())

	// TODO test
	// create a scope
	scopes.POST("", func(c *gin.Context) {
		newScope, valid := createScopesDto.Validate(c)
		if !valid {
			return
		}

		createScope.CreateScopes(c, scopeRepository, newScope)
	})

	//
	//// TODO test
	//// update a user
	//scopes.PUT("/:id", func(c *gin.Context) {
	//	userId, validated := objectIdParam.Validate(c)
	//	if !validated {
	//		return
	//	}
	//
	//	newUser, validated := updateUserDto.Validate(c, userCollection)
	//	if !validated {
	//		return
	//	}
	//
	//	updateUsers.UpdateUsers(c, scopeRepository, newUser, userId)
	//})
	//
	//// TODO test
	//// get one user
	//scopes.GET("/:id", func(c *gin.Context) {
	//	userId, validated := objectIdParam.Validate(c)
	//	if !validated {
	//		return
	//	}
	//
	//	findUsers.FindUsers(c, scopeRepository, userId)
	//})
	//
	//// TODO test
	//// get all scopes
	//scopes.GET("", func(c *gin.Context) {
	//	pag := pagination.Validate(c)
	//
	//	findAllUsers.FindAllUsers(c, scopeRepository, pag)
	//})

}
