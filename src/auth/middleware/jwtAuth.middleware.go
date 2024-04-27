package middleware

import (
	"gin-api/src/auth/repositories/authRepository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// JwtAuthMiddleware for jws
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authRepo := authRepository.MongoAuth{}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			authRepo.Unauthorize(c)
			c.Abort()
			return
		}

		// Parse token
		prefix := "Bearer "
		if !strings.HasPrefix(strings.ToUpper(tokenString), strings.ToUpper(prefix)) {
			authRepo.Unauthorize(c)
			c.Abort()
			return
		}

		// trim jwt
		jwtTokenString := tokenString[len(prefix):]

		jwtKey := os.Getenv("JWT_AUTH_SECRET_KEY")
		token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			authRepo.Unauthorize(c)
			c.Abort()
			return
		}

		// Extract user info from claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			authRepo.Unauthorize(c)
			c.Abort()
			return
		}

		// Access username from claims
		email, ok := claims["username"].(string)
		if !ok {
			authRepo.Unauthorize(c)
			c.Abort()
			return
		}

		// You can now use the username or any other user information as needed
		c.Set("email", email)
		// c.Get()

		// Token is valid, proceed
		c.Next()
	}
}
