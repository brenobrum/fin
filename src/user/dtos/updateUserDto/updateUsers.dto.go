package updateUserDto

import (
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	stringValidator "gin-api/domain/validators/string"
	"gin-api/src/user/constants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UpdateUsersDto struct {
	Name     *string `json:"name" bson:"name,omitempty"`
	Email    *string `json:"email" bson:"email,omitempty"`
	Password *string `json:"password" bson:"password,omitempty"`
}

func Validate(c *gin.Context, collection *mongo.Collection) (UpdateUsersDto, bool) {
	apiError := apiErros.NewApiError()

	var updateUserDto UpdateUsersDto
	if err := c.ShouldBindJSON(&updateUserDto); err != nil {
		apiError.SetMessage(http_exceptions.BAD_REQUEST)
		apiError.SetCode(http_exceptions.BAD_REQUEST_CODE)

		c.JSON(http.StatusBadRequest, apiError.Throw())
		return UpdateUsersDto{}, false
	}

	if updateUserDto.Name != nil {
		stringValidator.ValidateString(
			stringValidator.StringValidationOpts{FieldName: "name", Value: *updateUserDto.Name, ApiError: &apiError},
			stringValidator.IsNotEmpty,
			stringValidator.IsUnique(collection, "email"),
		)
	}

	if updateUserDto.Email != nil {
		stringValidator.ValidateString(
			stringValidator.StringValidationOpts{FieldName: "email", Value: *updateUserDto.Email, ApiError: &apiError},
			stringValidator.IsNotEmpty,
			stringValidator.IsEmail,
		)
	}

	if updateUserDto.Password != nil {
		stringValidator.ValidateString(
			stringValidator.StringValidationOpts{FieldName: "password", Value: *updateUserDto.Password, ApiError: &apiError},
			stringValidator.IsNotEmpty,
			stringValidator.MinLenght(constants.USER_PASSWORD_MIN_LENGTH),
		)
	}

	if apiError.HasErrors() {
		apiError.SetMessage(http_exceptions.BAD_REQUEST)
		apiError.SetCode(http_exceptions.BAD_REQUEST_CODE)

		c.JSON(http.StatusBadRequest, apiError.Throw())
		return UpdateUsersDto{}, false
	}

	return updateUserDto, true
}
