package writer

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_to_object_model_service/output/decorator"
)

// NewFileWriter constructs a new FileWriter.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{
		path: path,
		decorators: []decorator.Decorator{
			decorator.FormatDecorator{},
			decorator.ImportDecorator{},
		},
	}
}
