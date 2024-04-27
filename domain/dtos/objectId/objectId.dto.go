package objectIdParam

import (
	"gin-api/domain/types/apiErros"
	stringValidator "gin-api/domain/validators/string"
	"gin-api/domain/validators/string/isObjectId"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Validate(c *gin.Context) (primitive.ObjectID, bool) {
	apiError := apiErros.NewApiError()

	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	stringValidator.ValidateString(
		stringValidator.StringValidationOpts{FieldName: "id", Value: id, ApiError: &apiError},
		stringValidator.IsObjectId,
	)

	if apiError.HasErrors() {
		apiError.SetMessage(isObjectId.IS_OBJECT_ID_MESSAGE)
		apiError.SetCode(isObjectId.IS_OBJECT_ID_CODE)

		c.JSON(http.StatusBadRequest, apiError.Throw())
		return primitive.NilObjectID, false
	}

	return objId, true
}
