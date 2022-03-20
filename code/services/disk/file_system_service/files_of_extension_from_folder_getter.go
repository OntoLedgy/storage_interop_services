package file_system_service

import (
	"container/list"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/lists"
	"io/ioutil"
	"strings"
)

//https://github.com/boro-alpha/nf_common/blob/master/nf_common_source/code/services/file_system_service/files_of_extension_from_folder_getter.py

//import os
//from nf_common_source.code.services.file_system_service.objects.files import Files
//from nf_common_source.code.services.file_system_service.objects.folders import Folders

//def get_all_files_of_extension_from_folder(
func GetAllFilesOfExtensionFromFolder(
	//folder: Folders,
	folder *object_model.Folders,
	//dot_extension_string: str) \
	//-> list:
	dotExtensionString string) *lists.Lists {

	//list_of_files_of_extension = \
	listOfFilesOfExtension :=
		//list()
		&lists.Lists{
			new(list.List),
		}

	listOfFilesOfExtension.Init()

	filesInFolder, folderReadError := ioutil.ReadDir(folder.Path.String())

	if folderReadError != nil {
		panic("cannot read folder")
	}

	//for file_name in os.listdir(folder.absolute_path_string):
	for _, file := range filesInFolder {

		//list_of_files_of_extension = \
		listOfFilesOfExtension =
			//__add_file_of_specific_extension(
			addFileOfSpecificExtension(
				//file_name=file_name,
				file.Name(),
				//dot_extension_string=dot_extension_string,
				dotExtensionString,
				//list_of_files_of_extension=list_of_files_of_extension,
				listOfFilesOfExtension,
				//folder=folder)
				folder)
	}

	//return \
	//list_of_files_of_extension
	return listOfFilesOfExtension
}

//def __add_file_of_specific_extension(
func addFileOfSpecificExtension(
	//file_name: str,
	fileName string,
	//dot_extension_string: str,
	dotExtensionString string,
	//list_of_files_of_extension: list,
	listOfFilesOfExtension *lists.Lists,
	//folder: Folders) \
	//-> list:
	folder *object_model.Folders) *lists.Lists {

	//if not file_name.endswith(dot_extension_string):
	if !strings.HasSuffix(fileName,
		dotExtensionString) {
		//return list_of_files_of_extension

		return listOfFilesOfExtension

	}

	//file_absolute_path_string = \
	fileAbsolutePathString :=
		//os.path.join(
		folder.Path.Join(fileName)
	//folder.absolute_path_string,
	//file_name)

	//file = \
	file :=
		&object_model.Files{}
	//Files(
	file.Initialise(
		//absolute_path_string=file_absolute_path_string)
		fileAbsolutePathString.String())

	//list_of_files_of_extension.append(
	listOfFilesOfExtension.List.PushFront(
		//file)
		file)

	//return \
	//list_of_files_of_extension
	return listOfFilesOfExtension
}
