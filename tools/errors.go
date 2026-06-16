package tools

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func NewValidationError(field, msg string) *ValidationError {
	return &ValidationError{
		Errors: []FieldError{
			{Field: field, Message: msg},
		},
	}
}

func (v *ValidationError) Error() string {
	return "validation error"
}

func HandlerSimpleError(ctx *gin.Context, httpCode int, message string, err error) {
	if message != "" {
		ctx.JSON(httpCode, gin.H{"errors": fmt.Sprintf("%s", message)})
	} else {
		ctx.JSON(httpCode, gin.H{"errors": fmt.Sprintf("%s", err)})
	}
}

func HandleLogError(err error, message string) {
	log.Printf("%s: %s", message, err)
}
