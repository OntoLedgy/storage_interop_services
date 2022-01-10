package configurations

import (
	"fmt"
)

// These are the filename format command line parameter.

// FileNameFormat represents a output filename format

type FileNameFormat string

const (
	FileNameFormatCamelCase FileNameFormat = "c"
	FileNameFormatSnakeCase FileNameFormat = "s"
)

func (of FileNameFormat) String() string {
	return string(of)
}

// Set sets the datatype for the custom type for the flag package.
func (of *FileNameFormat) Set(s string) error {

	*of = FileNameFormat(s)
	if *of == "" {
		*of = FileNameFormatCamelCase
	}
	if !supportedFileNameFormats[*of] {
		return fmt.Errorf("filename format %q not supported", *of)
	}
	return nil
}
