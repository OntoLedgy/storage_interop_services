package configurations

import "fmt"

// These database_i_o_service types are supported.

// These are the output format command line parameter.
const (
	OutputFormatCamelCase OutputFormat = "c"
	OutputFormatOriginal  OutputFormat = "o"
)

// OutputFormat represents a output format option.
type OutputFormat string

// Set sets the datatype for the custom type for the flag package.
func (of *OutputFormat) Set(
	s string) error {
	*of = OutputFormat(s)
	if *of == "" {
		*of = OutputFormatCamelCase
	}
	if !supportedOutputFormats[*of] {
		return fmt.Errorf("output format %q not supported", *of)
	}
	return nil
}

// String is the implementation of the Stringer interface needed for flag.Value interface.
func (of OutputFormat) String() string {
	return string(of)
}
