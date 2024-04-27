package authRepository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	"gin-api/src/user/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type MongoAuth struct {
	Collection *mongo.Collection
}

func NewMongoAuthRepository(collection *mongo.Collection) MongoAuth {
	return MongoAuth{
		Collection: collection,
	}
}

func (m *MongoAuth) Auth(email, password string) bool {
	userFilter := bson.M{
		"email": email,
	}

	var user models.User
	// Perform the query to find one user
	if err := m.Collection.FindOne(context.Background(), userFilter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		} else {
			fmt.Println("Error decoding user:", err)
			return false
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func (m *MongoAuth) BasicAuth(c *gin.Context) (authenticated bool, email string) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return false, email
	}

	prefix := "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return false, email
	}

	encoded := strings.TrimPrefix(auth, prefix)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return false, email
	}

	pair := strings.SplitN(string(decoded), ":", 2)
	if len(pair) != 2 {
		return false, email
	}

	email = pair[0]
	password := pair[1]
	authenticated = m.Auth(email, password)

	return authenticated, email
}

func (m *MongoAuth) Unauthorize(c *gin.Context) {
	apiError := apiErros.NewApiError()

	apiError.SetMessage(http_exceptions.UNAUTHORIZED)
	apiError.SetCode(http_exceptions.UNAUTHORIZED_CODE)

	c.JSON(http.StatusUnauthorized, apiError.Throw())
}

func (m *MongoAuth) JWTAuth(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		m.Unauthorize(c)
		return
	}

	isAuthenticated := m.Auth(user.Email, user.Password)
	if !isAuthenticated {
		m.Unauthorize(c)
		return
	}

	tokenString := m.GenerateAuthToken(c, user.Email)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (m *MongoAuth) GenerateAuthToken(c *gin.Context, email string) string {
	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	authExp := os.Getenv("JWT_AUTH_EXP_MINUTES")
	authTimeInMinutes, _ := strconv.Atoi(authExp)

	claims["username"] = email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(authTimeInMinutes)).Unix() // set token exp

	// Sign the token with our secret
	jwtKey := os.Getenv("JWT_AUTH_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return ""
	}

	return tokenString
}
