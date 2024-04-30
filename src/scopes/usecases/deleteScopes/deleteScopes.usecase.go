package deleteScopes

import (
	"context"
	"gin-api/domain/types/apiErros"
	scopeRepositories "gin-api/src/scopes/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Exec(c *gin.Context, scopeRepository scopeRepositories.ScopeRepository, scopeId primitive.ObjectID) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	_, err := scopeRepository.Find(context.Background(), scopeId)

	// show error from database
	if apiError.InternalServerError(c, err) {
		return
	}

	deleted, err := scopeRepository.Delete(context.Background(), scopeId)
	// show error from database
	if apiError.InternalServerError(c, err) {
		return
	}

	if !deleted {
		c.Status(http.StatusNotFound)
		return
	}

	// return the user created
	c.Status(http.StatusOK)
}
