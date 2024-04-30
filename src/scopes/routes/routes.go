package routes

import (
	objectIdParam "gin-api/domain/dtos/objectId"
	"gin-api/domain/dtos/pagination"
	"gin-api/src/auth/middleware"
	createScopesDto "gin-api/src/scopes/dtos/createScopesDto"
	updateScopesDto "gin-api/src/scopes/dtos/updateScopesDto"
	scopeRepositories "gin-api/src/scopes/repositories"
	createScope "gin-api/src/scopes/usecases/createScopes"
	"gin-api/src/scopes/usecases/deleteScopes"
	"gin-api/src/scopes/usecases/findAllScopes"
	"gin-api/src/scopes/usecases/findScopes"
	"gin-api/src/scopes/usecases/updateScopes"
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

	// TODO test
	// update a scope
	scopes.PUT("/:id", func(c *gin.Context) {
		userId, validated := objectIdParam.Validate(c)
		if !validated {
			return
		}

		updatedScope, validated := updateScopesDto.Validate(c)
		if !validated {
			return
		}

		updateScopes.UpdateScope(c, scopeRepository, updatedScope, userId)
	})

	// TODO test
	// get one scope
	scopes.GET("/:id", func(c *gin.Context) {
		scopeId, validated := objectIdParam.Validate(c)
		if !validated {
			return
		}

		findScopes.FindScope(c, scopeRepository, scopeId)
	})

	// TODO test
	// get all scopes
	scopes.GET("", func(c *gin.Context) {
		pag := pagination.Validate(c)

		findAllScopes.FindAllScopes(c, scopeRepository, pag)
	})

	// TODO test
	// get all scopes
	scopes.DELETE("/:id", func(c *gin.Context) {
		scopeId, validated := objectIdParam.Validate(c)
		if !validated {
			return
		}

		deleteScopes.Exec(c, scopeRepository, scopeId)
	})

}
