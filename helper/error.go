package helper

import (
	"errors"
	"fmt"
)

var (
	ErrFailedValidation = errors.New("validation failed")

	ErrFailedValidationWrap = func(err error) error {
		return fmt.Errorf("%w: %v", ErrFailedValidation, err)
	}
)
