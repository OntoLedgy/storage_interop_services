package file_system_service

import (
	"fmt"
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

func SelectFile() *object_model.Files { //TODO: should be file

	filePath, _ := dialog.
		File().Title("Select File").Filter("Text", "txt").Load()

	file := &object_model.Files{}

	file.Initialise(filePath)

	return file

}

func SelectFolder() string { //TODO: should be Folders

	folderPath, file_selection_err :=
		dialog.Directory().Title("Select file storage location").Browse()

	if file_selection_err != nil {

		fmt.Println(
			file_selection_err)

		return ""

	} else {

		return folderPath
	}

}
