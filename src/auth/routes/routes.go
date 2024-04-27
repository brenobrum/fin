package routes

import (
	"gin-api/src/auth/refreshToken/repositories/refreshTokenRepository"
	"gin-api/src/auth/repositories/authRepository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func SetupAuthRoutes(router *gin.Engine, client *mongo.Client) {
	userCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	authRepo := authRepository.NewMongoAuthRepository(userCollection)
	refreshTokenRepo := refreshTokenRepository.NewMongoRefreshToken(client)

	auth := router.Group("/auth")

	auth.POST("/login", func(c *gin.Context) {
		authenticated, email := authRepo.BasicAuth(c)

		if !authenticated {
			authRepo.Unauthorize(c)
			return
		}

		accessJwtToken := authRepo.GenerateAuthToken(c, email)

		if len(accessJwtToken) == 0 {
			authRepo.Unauthorize(c)
			return
		}

		refreshToken := refreshTokenRepo.Set(email)
		if len(accessJwtToken) == 0 {
			authRepo.Unauthorize(c)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":  accessJwtToken,
			"refreshToken": refreshToken,
		})
		return
	})

	auth.POST("/refresh-token", func(c *gin.Context) {
		authenticated, email := refreshTokenRepo.Find(c)

		if !authenticated {
			authRepo.Unauthorize(c)
			return
		}

		accessJwtToken := authRepo.GenerateAuthToken(c, email)

		if len(accessJwtToken) == 0 {
			authRepo.Unauthorize(c)
			return
		}

		refreshToken := refreshTokenRepo.Set(email)
		if len(accessJwtToken) == 0 {
			authRepo.Unauthorize(c)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":  accessJwtToken,
			"refreshToken": refreshToken,
		})
		return
	})
}
