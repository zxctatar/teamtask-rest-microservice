package handlvalidator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	requiredTag = "required"
)

var (
	requiredTagErr = "field is required"
	defaultErr     = "field is invalid"
)

func MapValidationErrors(err error) (map[string]string, bool) {
	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		errMap := make(map[string]string)
		for _, e := range valErr {
			field := e.Field()
			tag := e.Tag()

			errMap[field] = validateError(tag)
		}
		return errMap, true
	}
	return nil, false
}

func validateError(tag string) string {
	switch tag {
	case requiredTag:
		return requiredTagErr
	default:
		return defaultErr
	}
}
