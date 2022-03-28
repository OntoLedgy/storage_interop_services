package object_model

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/wrappers"
	"io"
	"log"
	"os"
	"path/filepath"
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

func (fileSystemObject *FileSystemObjects) CreateIfNonExistent() (bool, error) {

	fileSystemObjectExists := fileSystemObject.Exists()

	if fileSystemObjectExists {
		return true, nil
	} else {

		folderCreationError := os.Mkdir(
			fileSystemObject.AbsolutePathString(),
			0755)

		if folderCreationError != nil {
			panic(folderCreationError)
		}
		return true, nil
	}

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

func (fileSystemObject *FileSystemObjects) Copy(targetPath string) (int64, error) {

	sourceFileStat, err := os.Stat(fileSystemObject.AbsolutePathString())
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf(
			"%s is not a regular file",
			fileSystemObject.AbsolutePathString())
	}

	source, err :=
		os.Open(
			fileSystemObject.AbsolutePathString())

	if err != nil {
		return 0, err
	}
	defer source.Close()

	directorCreationError := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)

	if directorCreationError != nil {
		log.Println(err)
	}

	destination, err := os.Create(targetPath)

	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)

	return nBytes, err
}
