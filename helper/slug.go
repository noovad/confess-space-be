package helper

import (
	"regexp"
	"strings"
)

var (
	nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9\s-]`)

	multipleSpacesRegex = regexp.MustCompile(`[\s-]+`)
)

func ToSlug(text string) string {
	text = strings.ToLower(text)

	text = nonAlphanumericRegex.ReplaceAllString(text, "")

	text = multipleSpacesRegex.ReplaceAllString(text, "-")

	return strings.Trim(text, "-")
}
