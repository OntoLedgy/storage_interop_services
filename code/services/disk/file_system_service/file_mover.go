package file_system_service

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/infrastructure/logging"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func MoveFiles(
	moveFileListFilename string,
	deleteSourceFlag string) (int, error) {

	moveFileListFile, csvData :=
		csv.OpenCsvFile(
			moveFileListFilename)

	moveFileList := csv.Read_csv_to_slice(
		moveFileListFile,
		csvData,
		",")

	var errorCount int = 0

	fmt.Printf("number of files to process : %s\n", len(moveFileList))

	log_file := logging.Set_log_file()

	defer log_file.Close()

	for index, row := range moveFileList {

		if index != 0 {

			log.Printf(
				"copying %s to %s \n", row[0], row[1])

			sourceFileStat, sourceFileStatsError :=
				os.Stat(row[0])

			if sourceFileStatsError != nil {
				log.Printf("source file error: %s\n", sourceFileStatsError)
				errorCount += 1
				continue
			}

			if !sourceFileStat.Mode().IsRegular() {

				log.Printf("%s",
					fmt.Errorf(
						"%s is not a regular file\n",
						row[0]))

				errorCount += 1
			}

			source, sourceFileOpenError := os.Open(row[0])

			if sourceFileOpenError != nil {

				log.Printf(
					"source file error %s\n",
					sourceFileOpenError)

				errorCount += 1

				continue
			}
			defer source.Close()

			destination_directory :=
				filepath.Dir(
					row[1])

			_, destinationDirectoryStatsError :=
				os.Stat(
					destination_directory)

			if os.IsNotExist(destinationDirectoryStatsError) {
				log.Printf("target directory %s does not exits, creating now\n", destination_directory)
				os.MkdirAll(destination_directory, os.ModePerm)

			}

			destination, destinationFileStatsError :=
				os.Create(
					row[1])

			if destinationFileStatsError != nil {
				log.Printf("destination file error: %s\n", sourceFileStatsError)
				errorCount += 1
				continue
			}

			defer destination.Close()

			if deleteSourceFlag == "yes" {
				fileMoveError := os.Rename(source.Name(), destination.Name())

				if fileMoveError != nil {
					log.Printf("cannot move file due to %v\n", fileMoveError)
					errorCount += 1
				} else {
					bytesCopied, fileCopyError := io.Copy(destination, source)

					if fileCopyError != nil {
						log.Printf("copied file error: %v\n", fileCopyError)
						errorCount += 1
					}
					log.Printf("sucessfully copied %v bytes\n", bytesCopied)
				}

			}

		}

	}
	log.Printf("Process completed with %v errors\n", errorCount)
	return errorCount, nil
}
