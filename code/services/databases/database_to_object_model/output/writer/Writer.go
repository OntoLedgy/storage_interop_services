package writer

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/output/decorator"
	"io/ioutil"
	"path"
)

const (
	// FileWriterExtension is the extension to write file_system_service of.
	FileWriterExtension = ".go"
)

// FileWriter is a writer that writes to a file given by the path and the table name.
type FileWriter struct {
	path       string
	decorators []decorator.Decorator
}

// Write is the implementation of the Writer interface. The FilerWriter writes
// decorated content to the file specified by the given path and table name.
func (w FileWriter) Write(
	tableName string,
	content string) error {

	fileName :=
		path.Join(
			w.path,
			tableName+FileWriterExtension)

	decorated, err :=
		w.decorate(
			content)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(decorated), 0666)
}

// decorate applies some decorations like formatting and empty import removal.
func (w FileWriter) decorate(
	content string) (
	decorated string,
	err error) {
	for _, fileDecorator := range w.decorators {
		content, err = fileDecorator.Decorate(content)
		if err != nil {
			return content, err
		}
	}

	return content, nil
}
