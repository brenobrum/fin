package findScopes

import (
	"context"
	"gin-api/domain/types/apiErros"
	scopeRepositories "gin-api/src/scopes/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func FindScope(c *gin.Context, findScope scopeRepositories.ScopeRepository, scopeId primitive.ObjectID) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	scope, err := findScope.Find(context.Background(), scopeId)

	// show error from database
	if apiError.InternalServerError(c, err) {
		return
	}

	// return the user created
	c.JSON(http.StatusOK, scope)
}
