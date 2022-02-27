package decorator

import (
	"strings"
)

// ImportDecorator removes empty import statements from the given content.
type ImportDecorator struct{}

// Decorate is the implementation of the Decorator interface.
func (ImportDecorator) Decorate(content string) (string, error) {
	// fight the symptom instead of the cause - if we didnt imported anything, remove it
	decorated := strings.ReplaceAll(content, "\nimport ()\n", "")
	return decorated, nil
}
