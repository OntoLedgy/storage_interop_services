package testing

import (
	"fmt"
	files "github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"testing"
)

func TestFileCreateReadDelete(t *testing.T) {

	fileNameAndPath :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\test.txt"

	fmt.Println("testing file")

	files.Create_file_if_does_not_exist_else_delete_it(fileNameAndPath)

	files.Delete_file_it_already_exists(fileNameAndPath)

}

func TestFileSelect(t *testing.T) {

	file := files.SelectFile()

	fmt.Println(file.BaseName())

}
