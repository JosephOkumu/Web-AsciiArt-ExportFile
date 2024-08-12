package web

import (
	"errors"
	"strings"
)

// ValidateInput validates the user input.
func ValidateInput(args []string) (string, error) {
	// Check if there is at least one argument (the input text)
	if len(args) < 2 {
		return "", errors.New("usage: go run . \"text to present as ascii art\"")
	}

	// Join all arguments except the program name and return as a single string
	return strings.Join(args[1:], " "), nil
}
