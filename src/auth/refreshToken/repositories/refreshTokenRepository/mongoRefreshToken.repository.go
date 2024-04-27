package refreshTokenRepository

import (
	"context"
	"errors"
	"fmt"
	"gin-api/domain/utils/randoString"
	"gin-api/src/auth/refreshToken/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"strconv"
	"strings"
	"time"
)

type MongoRefreshToken struct {
	collection *mongo.Collection
}

func NewMongoRefreshToken(client *mongo.Client) MongoRefreshToken {
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("refreshTokens")

	return MongoRefreshToken{collection: collection}
}

func (m *MongoRefreshToken) Set(email string) string {
	secureKey := randoString.RandoString(15)

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	authExp := os.Getenv("REFRESH_AUTH_EXP_MINUTES")
	authTimeInMinutes, _ := strconv.Atoi(authExp)

	claims["email"] = email
	claims["secureKey"] = secureKey
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(authTimeInMinutes)).Unix() // Token expires in 24 hours

	// Sign the token with your secret key
	secretKey := []byte(os.Getenv("REFRESH_AUTH_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	// Insert the document into the database
	document := bson.M{"email": email, "secureKey": secureKey}
	_, err = m.collection.InsertOne(context.Background(), document)
	if err != nil {
		return ""
	}

	return tokenString
}

func (m *MongoRefreshToken) Find(c *gin.Context) (bool, string) {

	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is not empty
	if authHeader == "" {
		return false, ""
	}

	// Split the Authorization header to get the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return false, ""
	}

	tokenString := parts[1]
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_AUTH_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Error parsing or validating token:", err)
		return false, ""
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error extracting claims from token")
		return false, ""
	}

	// Extract email and secure key from claims
	email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("Error extracting email from token claims")
		return false, ""
	}

	secureKey, ok := claims["secureKey"].(string)
	if !ok {
		fmt.Println("Error extracting secureKey from token claims")
		return false, ""
	}

	filter := bson.M{"email": email, "secureKey": secureKey}
	var result models.RefreshToken // Use an empty struct to check if any document matches the filter
	err = m.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Document not found
			return false, ""
		}

		// Other error occurred
		fmt.Println("Error finding document:", err)
		return false, ""
	}

	// Document found
	return true, email
}
