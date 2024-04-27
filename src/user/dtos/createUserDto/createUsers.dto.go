package createUserDto

import (
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	stringValidator "gin-api/domain/validators/string"
	"gin-api/src/user/constants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type CreateUsersDto struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func Validate(c *gin.Context, collection *mongo.Collection) (CreateUsersDto, bool) {
	apiError := apiErros.NewApiError()

	var userDto CreateUsersDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		apiError.SetMessage(http_exceptions.BAD_REQUEST)
		apiError.SetCode(http_exceptions.BAD_REQUEST_CODE)

		c.JSON(http.StatusBadRequest, apiError.Throw())
		return userDto, false
	}

	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "email", Value: userDto.Email, ApiError: &apiError},
		stringValidator.IsEmail,
		stringValidator.IsNotEmpty,
		stringValidator.IsUnique(collection, "email"),
	)
	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "name", Value: userDto.Name, ApiError: &apiError},
		stringValidator.IsNotEmpty,
	)
	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "password", Value: userDto.Password, ApiError: &apiError},
		stringValidator.IsNotEmpty,
		stringValidator.MinLenght(constants.USER_PASSWORD_MIN_LENGTH),
	)

	if apiError.HasErrors() {
		c.JSON(http.StatusBadRequest, apiError.Throw())
		return userDto, false
	}

	return userDto, true
}
