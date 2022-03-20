package file_system_service

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/sqweek/dialog"
)

func CreateFile() object_model.Files { //TODO: should be file

	filePath, _ := dialog.
		File().Title("Select File").Filter("Text", "txt").Save()

	file := object_model.Files{}

	file.Initialise(filePath)

	return file

}

func SelectFile(title string) *object_model.Files { //TODO: should be file

	filePath, _ := dialog.
		File().Title(title).Filter("Text", "txt").Load()

	file := &object_model.Files{}

	file.Initialise(filePath)

	return file

}
