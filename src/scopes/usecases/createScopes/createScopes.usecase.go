package createScope

import (
	"context"
	"fmt"
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	createScopesDto "gin-api/src/scopes/dtos/createUserDto"
	scopeRepositories "gin-api/src/scopes/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateScopes(c *gin.Context, mongoScope scopeRepositories.ScopeRepository, newScope createScopesDto.CreateScopesDto) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the scope data
	createdScope, err := mongoScope.Create(context.Background(), newScope)

	// show error from database
	if err != nil {
		fmt.Println(err)
		apiError.SetMessage(http_exceptions.INTERNAL_SERVER_ERROR)
		apiError.SetCode(http_exceptions.INTERNAL_SERVER_ERROR_CODE)
		apiError.SetDatailedError(err.Error())

		c.JSON(http.StatusInternalServerError, apiError.Throw())
		return
	}

	// return the user created
	c.JSON(http.StatusCreated, createdScope)
}
