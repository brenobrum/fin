package stringValidator

import (
	"gin-api/domain/exceptions/default_exceptions"
	"gin-api/domain/types/apiErros"
)

type StringValidationOpts struct {
	FieldName string
	Value     string
	ApiError  *apiErros.ApiError
}

type StringValidationFunc func(StringValidationOpts)

func ValidateString(validationOpts StringValidationOpts, validationFunctions ...StringValidationFunc) {
	for _, fn := range validationFunctions {
		fn(validationOpts)
	}

	if validationOpts.ApiError.HasErrors() {
		validationOpts.ApiError.SetMessage(default_exceptions.VALIDATION_ERROR_MESSAGE)
		validationOpts.ApiError.SetCode(default_exceptions.VALIDATION_ERROR_CODE)
	}
}
