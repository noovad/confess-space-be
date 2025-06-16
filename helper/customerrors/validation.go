package customerror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation = fmt.Errorf("Validation")
)

func WrapValidation(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return fmt.Errorf("%w: %s", ErrValidation, parseValidationErrors(ve))
	}
	return fmt.Errorf("%w: %v", ErrValidation, err)
}

func parseValidationErrors(errs validator.ValidationErrors) string {
	var messages []string
	for _, e := range errs {
		field := e.Field()
		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", field))
		case "uuid":
			messages = append(messages, fmt.Sprintf("%s must be a valid UUID", field))
		case "email":
			messages = append(messages, fmt.Sprintf("%s must be a valid email", field))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters", field, e.Param()))
		case "max":
			messages = append(messages, fmt.Sprintf("%s must be at most %s characters", field, e.Param()))
		default:
			messages = append(messages, fmt.Sprintf("%s is not valid", field))
		}
	}
	return strings.Join(messages, ", ")
}
