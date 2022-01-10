package configurations

import "fmt"

// These null types are supported. The types native and primitive map to the same
// underlying builtin golang type.
const (
	NullTypeSQL       NullType = "sql"
	NullTypeNative    NullType = "native"
	NullTypePrimitive NullType = "primitive"
)

// NullType represents a null type.
type NullType string

// SetCustomDatatype sets the datatype for the custom type for the flag package.
func (t *NullType) Set(
	s string) error {
	*t = NullType(s)

	if *t == "" {
		*t = NullTypeSQL
	}

	if !SupportedNullTypes[*t] {
		return fmt.Errorf(
			"null type %q not supported! supported: %v",
			*t,
			SprintfSupportedNullTypes())
	}
	return nil
}

// String is the implementation of the Stringer interface needed for flag.Value interface.
func (t NullType) String() string {
	return string(t)
}
