package apiErros

import (
	"gin-api/domain/types/apiErros/exceptions"
	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Message       string                     `json:"message"`
	Code          string                     `json:"code"`
	Errors        *map[string][]ErrorMessage `json:"errors,omitempty"`
	DetailedError *string                    `json:"detailedError,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewApiError() ApiError {
	return ApiError{
		Message: exceptions.TODO_MESSAGE,
		Code:    exceptions.TODO_ERROR,
	}
}

func (e *ApiError) Throw() gin.H {
	ginError := gin.H{
		"message": e.Message,
		"code":    e.Code,
	}

	if e.Errors != nil {
		ginError["errors"] = e.Errors
	}
	if e.DetailedError != nil {
		ginError["detailedError"] = e.Errors
	}

	return ginError
}

func (e *ApiError) SetMessage(message string) {
	e.Message = message
}

func (e *ApiError) SetCode(code string) {
	e.Code = code
}

func (e *ApiError) SetDatailedError(detailedError string) {
	e.DetailedError = &detailedError
}

func (e *ApiError) SetError(key string, error ErrorMessage) {
	if e.Errors == nil {
		errorsMap := make(map[string][]ErrorMessage)
		e.Errors = &errorsMap
	}

	// sets the error on the hash of erros
	(*e.Errors)[key] = append((*e.Errors)[key], error)
}

func (e *ApiError) HasErrors() bool {
	// validates erros length
	if e.Errors != nil && len(*e.Errors) != 0 {
		return true
	}

	return false
}
