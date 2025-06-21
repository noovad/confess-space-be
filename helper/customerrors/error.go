package customerror

import "errors"

var (
	ErrUserAlreadyInSpace = errors.New("user is already in the space")
)
