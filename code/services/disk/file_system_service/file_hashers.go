package file_system_service

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/infrastructure/logging"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"hash"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

//TODO: copy code from here. https://github.com/russtoku/parallel-copy-and-checksum/blob/master/psha1sum.go

//TODO: wrap in struct

func GetFileHashesForFolder(
	SourceFolder string, // TODO replace with Folders
	OutputFile string,
	hashAlgorithm string,
	skipFiles int,
	batchSize int) {
	var fileHash string

	var fileInformationTable [][]string
	var totalFileCount int

	filesList, directoryContents, fileWalkErrors := FindAllDirectoryContentRecursive(SourceFolder)

	if fileWalkErrors != nil {

		fmt.Println("error reading directory : %s", fileWalkErrors)

		return
	}
	logFile := logging.SetLogFile()

	defer logFile.Close()
	totalFileCount = len(filesList)
	sort.Strings(filesList)

	waitGroup := new(sync.WaitGroup)

	outputFile, _ :=
		csv.OpenCsvFile(
			OutputFile)

	fileInformationHeader := []string{
		"file_index",
		"file_name",
		"directory_full_path",
		"file_size",
		"file_modified_date",
		"file_mode",
		"item_is_directory",
		"file_hash_code"}

	fileInformationTable = append(fileInformationTable,
		fileInformationHeader)

	fileInformationTable = reportFolderContentsFileInformation(
		filesList,
		skipFiles,
		hashAlgorithm,
		directoryContents,
		totalFileCount,
		fileHash,
		waitGroup,
		fileInformationTable,
		batchSize,
		outputFile)

	fmt.Println("writing final : ", totalFileCount-(totalFileCount%batchSize), "-", totalFileCount, " / ", totalFileCount)
	csv.Write2dSliceSetToCsv(fileInformationTable, outputFile)

}

func reportFolderContentsFileInformation(
	filesList []string,
	skipFiles int,
	hashAlgorithm string,
	directoryContents []os.FileInfo,
	totalFileCount int,
	fileHash string,
	waitGroup *sync.WaitGroup,
	fileInformationTable [][]string,
	batchSize int,
	outputFile *os.File) [][]string {

	var fileInformationRow []string

	for fileIndex, file := range filesList {

		if fileIndex > skipFiles {

			fileInformationRow = reportFileInformation(
				fileIndex,
				hashAlgorithm,
				file,
				directoryContents,
				totalFileCount,
				fileHash,
				waitGroup)

			fileInformationTable = append(
				fileInformationTable,
				fileInformationRow)

			if fileIndex%batchSize == 0 {
				fmt.Println("writing : ",
					fileIndex-batchSize, "-",
					fileIndex, " / ",
					totalFileCount)

				csv.Write2dSliceSetToCsv(
					fileInformationTable,
					outputFile)

				fileInformationTable = nil
			}
		}
	}
	return fileInformationTable
}

func reportFileInformation(
	fileIndex int,
	hashAlgorithm string,
	file string,
	directoryContents []os.FileInfo,
	totalFileCount int,
	fileHash string,
	waitGroup *sync.WaitGroup) []string {

	log.Printf("calculting hash using %s for %s\r\n", hashAlgorithm, file)

	fileInformationRow := []string{
		strconv.Itoa(fileIndex),
		directoryContents[fileIndex].Name(),
		file,
		strconv.FormatInt(directoryContents[fileIndex].Size(), 10),
		directoryContents[fileIndex].ModTime().String(),
		directoryContents[fileIndex].Mode().String(),
		strconv.FormatBool(directoryContents[fileIndex].IsDir())}

	if directoryContents[fileIndex].IsDir() != true {

		fmt.Println("processing : ", fileIndex, " / ", totalFileCount)

		fileHash = CalculateFileHash(
			file,
			hashAlgorithm,
			waitGroup)

		fileInformationRow = append(
			fileInformationRow,
			fileHash)

	}
	return fileInformationRow

}

func CalculateFileHash(
	filePathAndName string, //TODO replace with Files
	hashAlgorithm string,
	waitGroup *sync.WaitGroup) string {

	file, err := os.Open(filePathAndName)

	if err != nil {
		panic(err.Error()) //TODO : add proper error handling (skip and log option)
	}
	var file_hash hash.Hash

	switch hashAlgorithm {
	case "sha256":
		file_hash = sha256.New()
	case "sha512":
		file_hash = sha512.New()
	case "md5":
		file_hash = md5.New()
	}

	// 2 channels: used to give green light for reading into buffer b1 or b2
	read_data_channel, read_status_channel := make(chan int, 1), make(chan int, 1)

	// 2 channels: used to give green light for hashing the content of b1 or b2
	hash_data_channel, hash_status_channel := make(chan int, 1), make(chan int, 1)

	// Start signal: Allow b1 to be read and hashed
	read_data_channel <- 1
	hash_data_channel <- 1

	waitGroup.Add(1)

	go hashHelper(
		file,
		file_hash,
		read_data_channel,
		read_status_channel,
		hash_data_channel,
		hash_status_channel,
		waitGroup)

	waitGroup.Add(1)

	hashHelper(
		file,
		file_hash,
		read_status_channel,
		read_data_channel,
		hash_status_channel,
		hash_data_channel,
		waitGroup)

	waitGroup.Wait()

	file_hash_code := fmt.Sprintf("%x", file_hash.Sum(nil))
	var file_hash_bytes []byte
	file_hash_code = string(b64.StdEncoding.EncodeToString(file_hash.Sum(file_hash_bytes)))
	fmt.Println("file path: ", filePathAndName, " --> hash code:", file_hash_code)
	return file_hash_code
}

func hashHelper(
	file *os.File,
	fileHash hash.Hash,
	mayRead <-chan int,
	readDone chan<- int,
	mayHash <-chan int,
	hashDone chan<- int,
	waitGroup *sync.WaitGroup) {

	for b, hasMore := make([]byte, 8192<<10), true; hasMore; {
		<-mayRead
		n, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				hasMore = false
			} else {
				panic(err) //TODO : add proper error handling (skip and log option)
			}
		}
		readDone <- 1

		<-mayHash
		_, err = fileHash.Write(b[:n])
		if err != nil {
			panic(err) //TODO : add proper error handling (skip and log option)
		}
		hashDone <- 1

	}
	waitGroup.Done()

}

//TODO : check if this is needed
func monitorWorker(wg *sync.WaitGroup, cs chan string) {
	wg.Wait()
	close(cs)
}
