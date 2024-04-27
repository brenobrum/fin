package updateUsers

import (
	"context"
	"fmt"
	"gin-api/domain/exceptions/http_exceptions"
	"gin-api/domain/types/apiErros"
	"gin-api/src/user/dtos/updateUserDto"
	"gin-api/src/user/models"
	"gin-api/src/user/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateUsers(c *gin.Context, mongoUser userRepositories.UserRepository, userUpdate updateUserDto.UpdateUsersDto, id primitive.ObjectID) {
	apiError := apiErros.NewApiError()

	// Call the usecase function, pass the user data
	updatedUser, err := mongoUser.Update(context.Background(), userUpdate, id)

	// show error from database
	if err != nil {
		fmt.Println(err)
		apiError.SetMessage(http_exceptions.INTERNAL_SERVER_ERROR)
		apiError.SetCode(http_exceptions.INTERNAL_SERVER_ERROR_CODE)
		apiError.SetDatailedError(err.Error())

		c.JSON(http.StatusInternalServerError, apiError.Throw())
		return
	}

	// return the user updated
	c.JSON(http.StatusOK, models.ResponseUser(updatedUser))
}
