package decorator

// Decorator represents an interface to decorate the given content.
type Decorator interface {
	Decorate(content string) (string, error)
}
