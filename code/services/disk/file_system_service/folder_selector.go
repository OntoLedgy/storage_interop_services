package file_system_service

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/wrappers"
	"github.com/sqweek/dialog"
)

func SelectFolder(title string) *object_model.Folders { //TODO: should be Folders

	folderPath, file_selection_err :=
		dialog.Directory().Title(title).Browse()

	absolutefolderPath := &wrappers.AbsolutePathWrappers{}
	absolutefolderPath.Initialise(folderPath)

	if file_selection_err != nil {

		fmt.Println(
			file_selection_err)

		return nil

	} else {
		folder := &object_model.Folders{}
		folder.Initialise(folderPath, nil)
		//folder.Path = absolutefolderPath

		return folder
	}

}
