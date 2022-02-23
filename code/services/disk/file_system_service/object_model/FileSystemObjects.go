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
func (f *FileSystemObjects) Initialise(
	//absolute_path_string: str):
	absolutePathString string) {

	var err error

	f.Uuid = &uuid_service.UUIDs{}

	//self.uuid = \
	f.Uuid, err =
		//create_new_uuid()
		uuid_service.GetUUID(
			1,  //TODO - add enumeration
			"") //TODO - wrap this inside the factory

	if err != nil {
		panic(err)
	}

	//self.__path = \
	f.Path =
		&wrappers.AbsolutePathWrappers{}

	//AbsolutePathWrappers(
	//absolute_path_string)
	f.Path.Initialise(
		absolutePathString)
}

//@property
//def base_name(
//self) \
//-> str:
func (f *FileSystemObjects) BaseName() string {
	//return \
	//self.__path.base_name
	return f.Path.BaseName()
}

//@property
//def absolute_path_string(
//self) \
//-> str:
func (f *FileSystemObjects) AbsolutePathString() string {
	//return \
	//self.__path.absolute_path_string
	return f.Path.AbsolutePathString()
}

//@property
//def absolute_level(
//self) \
//-> int:
func (f *FileSystemObjects) AbsoluteLevel() int {
	//return \
	//self.__path.absolute_level
	return f.Path.AbsoluteLevel()
}

//@property
//def parent_absolute_path_string(
//self) \
//-> str:
func (f *FileSystemObjects) ParentAbsolutePathString() string {
	//return \
	//str(self.__path.parent)
	return f.Path.Parent().PathString()
}

//def extend_path(
//self,
//path_extension: str) \
//-> str:
func (f *FileSystemObjects) ExtendPath(pathExtension string) string {
	//return \
	//self.__path.extend_path(
	//path_extension)

	return f.Path.ExtendPath(pathExtension)
}

//def exists(
//self) \
//-> bool:
func (f *FileSystemObjects) Exists() bool {
	//return \
	//self.__path.exists()
	return f.Path.Exists()
}

//def list_of_components(self):
func (f *FileSystemObjects) ListOfComponents() []string {
	//return \
	//self.__path.list_of_components()
	return f.Path.ListOfComponents()
}

//def item_count(
//self) \
//-> int:
func (f *FileSystemObjects) ItemCount() int {
	//return \
	//self.__path.item_count()
	return f.Path.ItemCount()
}
