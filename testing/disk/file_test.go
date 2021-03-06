package disk

import (
	"fmt"
	files "github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service/object_model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileCreateReadDelete(t *testing.T) {

	fileNameAndPath :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\test.txt"

	fmt.Println("testing file")

	files.CreateFileIfDoesNotExistElseDeleteIt(fileNameAndPath)

	files.DeleteFileIfItAlreadyExists(fileNameAndPath)

}

func TestFileSelect(t *testing.T) {

	file := files.SelectFile("select a file")

	fmt.Println(file.BaseName())

}

func TestFolderSelect(t *testing.T) {

	folder := files.SelectFolder("select a folder")

	fmt.Printf("basename: %s, full path : %s\n",
		folder.BaseName(),
		folder.Path)

}

func TestSelectFilesFromFolder(t *testing.T) {

	folder := files.SelectFolder("select a folder")

	list_of_files := files.GetAllFilesOfExtensionFromFolder(folder, "csv")

	fmt.Printf("list of files :\n %v", list_of_files.List.Len())

}

func TestCreateFolder(t *testing.T) {

	folderPath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\outputs\\testing_foldercreation2"

	folder := &object_model.Folders{}

	folder.Initialise(folderPath, nil)

	folder.CreateIfNonExistent()

	folderExists := folder.Exists()

	assert.True(t, folderExists)
}

func TestCopy(t *testing.T) {

	filePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\addresses.csv"

	targetFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\output\\addresses.csv"

	file := &object_model.Files{}

	file.Initialise(filePath, nil)

	file.CreateIfNonExistent()

	bytesCopied, fileCopied :=
		file.Copy(targetFilePath)

	fmt.Printf("bytes copied : %v", bytesCopied)

	assert.Nil(t, fileCopied)
}
