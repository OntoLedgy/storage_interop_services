package decorator

import (
	"fmt"
	"go/format"
)

// FormatDecorator applies a formatting decoration to the given content.
type FormatDecorator struct{}

// Decorate is the implementation of the Decorator interface.
func (FormatDecorator) Decorate(content string) (string, error) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		return content, fmt.Errorf("could not format content: %v", err)
	}
	return string(formatted), nil
}
