package createScopesDto

import (
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	stringValidator "gin-api/domain/validators/string"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateScopesDto struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

func Validate(c *gin.Context) (CreateScopesDto, bool) {
	apiError := apiErros.NewApiError()

	var scopeDto CreateScopesDto

	if err := c.ShouldBindJSON(&scopeDto); err != nil {
		apiError.SetMessage(http_exceptions.BAD_REQUEST)
		apiError.SetCode(http_exceptions.BAD_REQUEST_CODE)

		c.JSON(http.StatusBadRequest, apiError.Throw())
		return scopeDto, false
	}

	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "name", Value: scopeDto.Name, ApiError: &apiError},
		stringValidator.IsNotEmpty,
	)
	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "description", Value: scopeDto.Description, ApiError: &apiError},
	)

	if apiError.HasErrors() {
		c.JSON(http.StatusBadRequest, apiError.Throw())
		return scopeDto, false
	}

	return scopeDto, true
}
