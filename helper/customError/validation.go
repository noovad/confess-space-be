package customerror

import (
	"fmt"
)

var (
	ErrValidation = fmt.Errorf("validation failed")
)

func WrapValidation(err error) error {
	return fmt.Errorf("%w: %v", ErrValidation, err)
}
