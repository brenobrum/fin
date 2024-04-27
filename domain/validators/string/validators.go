package stringValidator

import (
	"fmt"
	"gin-api/domain/types/apiErros"
	"gin-api/domain/validators/string/isEmail"
	"gin-api/domain/validators/string/isNotEmpty"
	"gin-api/domain/validators/string/isObjectId"
	"gin-api/domain/validators/string/isUnique"
	"gin-api/domain/validators/string/maxLength"
	"gin-api/domain/validators/string/minLength"
	"go.mongodb.org/mongo-driver/mongo"
)

// IsEmail validator.
func IsEmail(validationOpts StringValidationOpts) {
	if !isEmail.IsEmail(validationOpts.Value) {
		validationOpts.ApiError.SetError(validationOpts.FieldName, apiErros.ErrorMessage{Message: isEmail.IS_NOT_EMAIL_MESSAGE, Code: isEmail.IS_NOT_EMAIL_CODE})
	}
}

// IsNotEmpty validator
func IsNotEmpty(validationOpts StringValidationOpts) {
	if !isNotEmpty.IsNotEmpty(validationOpts.Value) {
		validationOpts.ApiError.SetError(validationOpts.FieldName, apiErros.ErrorMessage{Message: isNotEmpty.IS_EMPTY_MESSAGE, Code: isNotEmpty.IS_EMPTY_CODE})
	}
}

// MaxLenght validator
func MaxLenght(n int) StringValidationFunc {

	return func(validationOptions StringValidationOpts) {
		if !maxLength.MaxLength(validationOptions.Value, n) {

			message := fmt.Sprintf("%s %d characters", maxLength.MAX_LENGHT_MESSAGE, n)
			validationOptions.ApiError.SetError(validationOptions.FieldName, apiErros.ErrorMessage{Message: message, Code: maxLength.MAX_LENGHT_CODE})
		}
	}
}

// MinLenght validator
func MinLenght(n int) StringValidationFunc {

	return func(validationOptions StringValidationOpts) {
		if !minLength.MinLength(validationOptions.Value, n) {

			message := fmt.Sprintf("%s %d characters", minLength.MIN_LENGHT_MESSAGE, n)
			validationOptions.ApiError.SetError(validationOptions.FieldName, apiErros.ErrorMessage{Message: message, Code: minLength.MIN_LENGHT_CODE})
		}
	}
}

// IsObjectId validator
func IsObjectId(validationOpts StringValidationOpts) {
	if !isObjectId.IsObjectId(validationOpts.Value) {
		validationOpts.ApiError.SetError(validationOpts.FieldName, apiErros.ErrorMessage{Message: isObjectId.IS_OBJECT_ID_MESSAGE, Code: isObjectId.IS_OBJECT_ID_CODE})
	}
}

// IsUnique validator
func IsUnique(mongoCollection *mongo.Collection, attributeName string) StringValidationFunc {

	return func(validationOptions StringValidationOpts) {
		if !isUnique.IsUnique(mongoCollection, attributeName, validationOptions.Value) {

			validationOptions.ApiError.SetError(validationOptions.FieldName, apiErros.ErrorMessage{Message: isUnique.IS_NOT_UNIQUE_MESSAGE, Code: isUnique.IS_NOT_UNIQUE_CODE})
		}
	}
}
