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

func Get_file_hashes_for_folder(
	SourceFolder string, // TODO replace with Folders
	OutputFile string,
	hashAlgorithm string,
	skipFiles int,
	batch_size int) {
	var file_hash string

	var file_information_row []string
	var file_information_table [][]string
	var total_file_count int

	files_list, directory_contents, file_walk_errors := Find_all_directory_content_recursive(SourceFolder)

	if file_walk_errors != nil {

		fmt.Println("error reading directory : %s", file_walk_errors)

		return
	}
	log_file := logging.Set_log_file()

	defer log_file.Close()
	total_file_count = len(files_list)
	sort.Strings(files_list)

	wait_group := new(sync.WaitGroup)

	output_file, _ :=
		csv.Open_csv_file(
			OutputFile)

	for file_index, file := range files_list {

		if file_index > skipFiles {

			log.Printf("calculting hash using %s for %s\r\n", hashAlgorithm, file)
			file_information_row = nil
			file_information_row = append(
				file_information_row,
				strconv.Itoa(file_index),
				directory_contents[file_index].Name(),
				file,
				strconv.FormatInt(directory_contents[file_index].Size(), 10),
				directory_contents[file_index].ModTime().String(),
				directory_contents[file_index].Mode().String(),
				strconv.FormatBool(directory_contents[file_index].IsDir()))

			if directory_contents[file_index].IsDir() != true {

				fmt.Println("processing : ", file_index, " / ", total_file_count)

				file_hash = Calculate_file_hash(
					file,
					hashAlgorithm,
					wait_group)
				file_information_row = append(
					file_information_row,
					file_hash)

			}

			file_information_table = append(file_information_table, file_information_row)

			if file_index%batch_size == 0 {
				fmt.Println("writing : ", file_index-batch_size, "-", file_index, " / ", total_file_count)
				csv.Write_2d_slice_set_to_csv(file_information_table, output_file)
				file_information_table = nil
			}

		}
	}

	fmt.Println("writing final : ", total_file_count-(total_file_count%batch_size), "-", total_file_count, " / ", total_file_count)
	csv.Write_2d_slice_set_to_csv(file_information_table, output_file)

}

func monitorWorker(wg *sync.WaitGroup, cs chan string) {
	wg.Wait()
	close(cs)
}

func Calculate_file_hash(
	file_path_and_name string, //TODO replace with Files
	hash_algorithm string,
	wait_group *sync.WaitGroup) string {

	file, err := os.Open(file_path_and_name)

	if err != nil {
		panic(err.Error()) //TODO : add proper error handling (skip and log option)
	}
	var file_hash hash.Hash

	switch hash_algorithm {
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

	wait_group.Add(1)

	go hashHelper(
		file,
		file_hash,
		read_data_channel,
		read_status_channel,
		hash_data_channel,
		hash_status_channel,
		wait_group)

	wait_group.Add(1)

	hashHelper(
		file,
		file_hash,
		read_status_channel,
		read_data_channel,
		hash_status_channel,
		hash_data_channel,
		wait_group)

	wait_group.Wait()

	file_hash_code := fmt.Sprintf("%x", file_hash.Sum(nil))
	var file_hash_bytes []byte
	file_hash_code = string(b64.StdEncoding.EncodeToString(file_hash.Sum(file_hash_bytes)))
	fmt.Println("file path: ", file_path_and_name, " --> hash code:", file_hash_code)
	return file_hash_code
}

func hashHelper(
	file *os.File,
	file_hash hash.Hash,
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
		_, err = file_hash.Write(b[:n])
		if err != nil {
			panic(err) //TODO : add proper error handling (skip and log option)
		}
		hashDone <- 1

	}
	waitGroup.Done()

}
