package file_system_service

import (
	"fmt"
	"os"
)

func CreateFileIfDoesNotExistElseDeleteIt(filepath string) {
	// check if file exists
	var _, err = os.Stat(filepath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(filepath)
		if isError(err) {
			return
		}
		defer file.Close()
		fmt.Println("New file created Successfully", filepath)

	} else {
		var err = os.Remove(filepath)
		if isError(err) {
			return
		}
		fmt.Println("Existing file Deleted")
	}

}

func DeleteFileIfItAlreadyExists(filepath string) {

	var _, err = os.Stat(filepath)

	file_already_exists := !os.IsNotExist(err)

	if file_already_exists {
		var err = os.Remove(filepath)
		if isError(err) {
			return
		}
		fmt.
			Printf("Existing file: %s Deleted\n", filepath)

	}

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
