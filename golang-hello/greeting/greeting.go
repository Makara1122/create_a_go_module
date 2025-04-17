package greeting

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Check if the name is empty and return an error if it is.
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message,	nil
}