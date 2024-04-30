package findAllScopes

import (
	"context"
	"gin-api/domain/dtos/pagination"
	"gin-api/domain/types/apiErros"
	scopeRepositories "gin-api/src/scopes/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllScopes(c *gin.Context, scopeRepository scopeRepositories.ScopeRepository, pag pagination.Pagination) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	scopes, err := scopeRepository.FindMany(context.Background(), pag)

	// show error from database
	if apiError.InternalServerError(c, err) {
		return
	}

	// return the user created
	c.JSON(http.StatusOK, pag.PaginatedInfo(scopes))
}
