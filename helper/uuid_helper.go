package helper

import (
	"errors"

	"github.com/google/uuid"
)

// StringToUUID converts a string to a UUID and validates its format
func StringToUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.Nil, errors.New("empty ID string")
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format")
	}

	return parsedUUID, nil
}

// IsValidUUID checks if a string is a valid UUID without converting it
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
