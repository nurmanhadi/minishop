package exception

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ProductNotFound        = errors.New("product not found")
	ProductValidationError = errors.New(catchErr)
)

var catchErr string

func ValidationError(err error) {
	var errors string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrors {
			errors = fmt.Sprintf("field '%s' failed on '%s'", validationErr.Field(), validationErr.Tag())
		}
	}
	catchErr = errors
}
