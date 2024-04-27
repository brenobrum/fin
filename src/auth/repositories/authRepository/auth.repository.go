package authRepository

import "github.com/gin-gonic/gin"

type AuthRepository interface {
	BasicAuth(c *gin.Context) bool
	JWTAuth(c *gin.Context) bool
	RefreshToken(c *gin.Context) string
	Unauthorize(c *gin.Context)
	GenerateAuthToken(c *gin.Context, email string) string
}
