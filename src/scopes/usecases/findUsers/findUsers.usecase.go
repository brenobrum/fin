package findUsers

import (
	"context"
	"gin-api/src/user/models"
	"gin-api/src/user/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func FindUsers(c *gin.Context, mongoUser userRepositories.UserRepository, userId primitive.ObjectID) {
	// Call the usecase function, pass the user data
	user, err := mongoUser.Find(context.Background(), userId)

	// show error from database
	if err != nil && err.Error() != "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.Status(http.StatusNotFound)
		return
	}

	// return the user created
	c.JSON(http.StatusOK, models.ResponseUser(*user))
}
