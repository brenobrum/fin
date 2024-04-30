package updateScopes

import (
	"context"
	"gin-api/domain/types/apiErros"
	updateScopesDto "gin-api/src/scopes/dtos/updateScopesDto"
	scopeRepositories "gin-api/src/scopes/repositories"
	_ "gin-api/src/user/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateScope(c *gin.Context, mongoScope scopeRepositories.ScopeRepository, scopeUpdate updateScopesDto.UpdateScopesDto, id primitive.ObjectID) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	updateScope, err := mongoScope.Update(context.Background(), scopeUpdate, id)

	// show error from database
	if apiError.InternalServerError(c, err) {
		return
	}

	// return the user updated
	c.JSON(http.StatusOK, updateScope)
}
