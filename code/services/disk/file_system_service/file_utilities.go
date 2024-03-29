package file_system_service

import (
	"os"
	"path/filepath"
)

func FindAllDirectoryContentRecursive(root string) ([]string, []os.FileInfo, error) {

	var files []string
	var files_information []os.FileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		files = append(files, path)
		files_information = append(files_information, info)

		return nil
	})
	return files, files_information, err
}
