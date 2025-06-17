package helper

import (
	"regexp"
	"strings"
)

var (
	// Match any character that is not alphanumeric, space, or hyphen
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9\s-]`)

	// Match multiple spaces or hyphens
	multipleSpacesRegex = regexp.MustCompile(`[\s-]+`)
)

// ToSlug converts a string to a URL-friendly slug
// Example: "Hello World!" -> "hello-world"
func ToSlug(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Replace non-alphanumeric characters with empty string
	text = nonAlphanumericRegex.ReplaceAllString(text, "")

	// Replace spaces and multiple hyphens with a single hyphen
	text = multipleSpacesRegex.ReplaceAllString(text, "-")

	// Trim hyphens from beginning and end
	return strings.Trim(text, "-")
}
