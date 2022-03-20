package object_model

import (
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/wrappers"
)

type FileSystemObjects struct {
	Uuid *uuid_service.UUIDs
	Path *wrappers.AbsolutePathWrappers
}

//def __init__(
//self,
func (fileSystemObject *FileSystemObjects) Initialise(
	//absolute_path_string: str):
	absolutePathString string) {

	var err error

	fileSystemObject.Uuid = &uuid_service.UUIDs{}

	//self.uuid = \
	fileSystemObject.Uuid, err =
		//create_new_uuid()
		uuid_service.GetUUID(
			1,  //TODO - add enumeration
			"") //TODO - wrap this inside the factory

	if err != nil {
		panic(err)
	}

	//self.__path = \
	fileSystemObject.Path =
		&wrappers.AbsolutePathWrappers{}

	//AbsolutePathWrappers(
	//absolute_path_string)
	fileSystemObject.Path.Initialise(
		absolutePathString)
}

//@property
//def base_name(
//self) \
//-> str:
func (fileSystemObject *FileSystemObjects) BaseName() string {
	//return \
	//self.__path.base_name
	return fileSystemObject.Path.BaseName()
}

//@property
//def absolute_path_string(
//self) \
//-> str:
func (fileSystemObject *FileSystemObjects) AbsolutePathString() string {
	//return \
	//self.__path.absolute_path_string
	return fileSystemObject.Path.AbsolutePathString()
}

//@property
//def absolute_level(
//self) \
//-> int:
func (fileSystemObject *FileSystemObjects) AbsoluteLevel() int {
	//return \
	//self.__path.absolute_level
	return fileSystemObject.Path.AbsoluteLevel()
}

//@property
//def parent_absolute_path_string(
//self) \
//-> str:
func (fileSystemObject *FileSystemObjects) ParentAbsolutePathString() string {
	//return \
	//str(self.__path.parent)
	return fileSystemObject.Path.Parent().PathString()
}

//def extend_path(
//self,
//path_extension: str) \
//-> str:
func (fileSystemObject *FileSystemObjects) ExtendPath(pathExtension string) string {
	//return \
	//self.__path.extend_path(
	//path_extension)

	return fileSystemObject.Path.ExtendPath(pathExtension)
}

//def exists(
//self) \
//-> bool:
func (fileSystemObject *FileSystemObjects) Exists() bool {
	//return \
	//self.__path.exists()
	return fileSystemObject.Path.Exists()
}

//def list_of_components(self):
func (fileSystemObject *FileSystemObjects) ListOfComponents() []string {
	//return \
	//self.__path.list_of_components()
	return fileSystemObject.Path.ListOfComponents()
}

//def item_count(
//self) \
//-> int:
func (fileSystemObject *FileSystemObjects) ItemCount() int {
	//return \
	//self.__path.item_count()
	return fileSystemObject.Path.ItemCount()
}
