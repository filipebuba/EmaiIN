package internalError

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateStruct(intObj interface{}) error {
	validate := validator.New()
	err := validate.Struct(intObj)
	if err == nil {
		return nil
	}
	var validatorErrors validator.ValidationErrors
	errors.As(err, &validatorErrors)
	validationError := validatorErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationError.Param())
	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
