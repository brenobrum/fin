package createUsers

import (
	"context"
	"fmt"
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	"gin-api/src/user/dtos/createUserDto"
	"gin-api/src/user/models"
	userRepositories "gin-api/src/user/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUsers(c *gin.Context, mongoUser userRepositories.UserRepository, newUser createUserDto.CreateUsersDto) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	createdUser, err := mongoUser.Create(context.Background(), newUser)

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
	c.JSON(http.StatusCreated, models.ResponseUser(*createdUser))
}
