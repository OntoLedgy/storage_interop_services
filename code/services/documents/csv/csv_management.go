package csv

import (
	"encoding/csv"
	"fmt"

	"github.com/TomOnTime/utfutil"
	"github.com/saintfish/chardet"
	"io"
	"log"
	"os"
	"strings"
)

func Open_csv_file(csv_filename string) (*os.File, []byte) {

	csv_file, err :=
		os.OpenFile(
			csv_filename,
			os.O_CREATE|os.O_RDWR,
			0777)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// charset detection
	buffer := make([]byte, 32<<10)
	size, _ := io.ReadFull(csv_file, buffer)
	input := buffer[:size]

	var detector = chardet.NewTextDetector()
	result, err := detector.DetectBest(input)

	fmt.
		Printf(
			"Opening File: %s, File Ecoding : %s , Language: %s \n",
			csv_filename,
			result.Charset,
			result.Language)

	//#TODO - ADD Switch here based on the the encoding detected.

	var csv_data_utfutls []byte

	switch result.Charset {

	case "UTF-8":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF8)
	case "UTF-16LE":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF16LE)
	case "ISO-8859-1":
		csv_data_utfutls, _ = utfutil.ReadFile(csv_filename, utfutil.UTF8)
	}

	return csv_file, csv_data_utfutls

}

func Write_2d_slice_set_to_csv(
	slice_to_write [][]string,
	csv_file *os.File) bool {

	for _, slice_row := range slice_to_write {

		Write_1d_slice_to_csv(
			slice_row,
			csv_file)

	}

	return true
}

func Write_1d_slice_to_csv(
	slice_to_write []string,
	csv_file *os.File) bool {

	var slice_to_write_with_quotes = make([]string, len(slice_to_write))
	writer_to_file :=
		csv.NewWriter(csv_file)
	//#TODO add helper function to change []interface to []string (see https://stackoverflow.com/questions/33357156/write-struct-to-csv-file)

	for column_index, cell_value := range slice_to_write {

		slice_to_write_with_quotes[column_index] = cell_value
	}

	writer_to_file.Write(
		slice_to_write_with_quotes)

	writer_to_file.Flush()

	return true

}

func ReadCsvToSlice(csv_file_name string, delimiter string) [][]string {

	csv_file,
		csv_file_data :=
		Open_csv_file(
			csv_file_name)

	csv_dataset := Read_csv_to_slice(
		csv_file,
		csv_file_data,
		delimiter)

	return csv_dataset

}

func Read_csv_to_slice(csv_file *os.File, csv_data_utfutls []byte, delimiter string) [][]string {

	//--END OF UTFUTIL reader
	/***********easycsv reader
	easy_csv_reader := easycsv.NewReaderFile(file_name ,  easycsv.Option{
		Comma: '\t',
	})
	var rawCSVdata [][]string
	easy_csv_reader.ReadAll(&rawCSVdata)
	* END OF EASY CSV REader*/

	// csv reader standard

	csv_reader := csv.NewReader(strings.NewReader(string(csv_data_utfutls)))
	//csv_reader.TrimLeadingSpace = true
	csv_reader.FieldsPerRecord = -1 // see the Reader struct information below

	switch delimiter {
	case "tab":
		csv_reader.Comma = '\t'
	case "":

	}

	csv_reader.LazyQuotes = true
	csv_data, csv_reader_error := csv_reader.ReadAll()

	if csv_reader_error != nil {
		panic(csv_reader_error)
	}
	//--- END OF CSV READER
	fmt.
		Printf("Read csv data from: %s, length: %v\n", csv_file.Name(), len(csv_data))

	//raw csv data generation  (NOT USED) -
	/*/#TODO - move this out.
	rawCSVdata_bytes := make([][]byte, len(csv_data)*len(csv_data[0]))
	for _,raw_csv_data_row := range csv_data {
		for _, raw_csv_data_column := range raw_csv_data_row {

			rawCSVdata_bytes = append (rawCSVdata_bytes, []byte(raw_csv_data_column))

		}
	}

	//  -- end of raw data bytes generation*/

	return csv_data
}

func Make_dynamic_2d_byte_slice(cols int, rows int) [][]byte {

	var mat = make([][]byte, rows)

	for i := range mat {

		mat[i] = make([]byte, cols)
	}

	return mat
}

func Write_slice_to_csv_split_by_column(column_index int, csv_file *os.File) {

	r := csv.NewReader(csv_file)
	partitions := make(map[string][][]string)

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil

				save_partitions(partitions)
				return
			}

			log.Fatal(err)
		}

		process(column_index, rec, partitions)
	}
}

func Write_slice_with_header_to_csv(slice_to_write [][]string, slice_header []string, csv_file_name string) {

	output_csv_file, _ :=
		Open_csv_file(
			csv_file_name)

	Write_1d_slice_to_csv(slice_header, output_csv_file)
	Write_2d_slice_set_to_csv(slice_to_write, output_csv_file)

}

// prints only
func save_partitions(partitions map[string][][]string) {

	var file_name string

	for part, recs := range partitions {
		fmt.Println(part)
		file_name = `./test/` + part + `.csv`
		csv_file, _ := Open_csv_file(file_name)

		Write_2d_slice_set_to_csv(recs, csv_file)

		for _, rec := range recs {
			fmt.Println(rec)
		}
	}
}

// this can also write/append directly to a file
func process(col_index int, rec []string, partitions map[string][][]string) {
	//l := len(rec)
	part := rec[col_index]
	if p, ok := partitions[part]; ok {
		partitions[part] = append(p, rec)
	} else {
		partitions[part] = [][]string{rec}
	}
}

func Get_csv_with_headers(csv_data_with_headers [][]interface{}) []map[string]interface{} {

	var rows []map[string]interface{}

	var header []interface{}

	for index, csv_data_with_headers_row := range csv_data_with_headers {

		if index == 0 {
			header = csv_data_with_headers_row
		} else {

			dict := make(map[string]interface{})

			for i := range header {
				dict[header[i].(string)] = csv_data_with_headers_row[i]
			}
			rows = append(rows, dict)

		}

	}

	return rows

}
