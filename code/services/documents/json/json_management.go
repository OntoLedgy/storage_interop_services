package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Read_json_to_byte_array(
	json_file_name string) []byte {

	json_data, json_file_read_error :=
		os.Open(
			json_file_name)

	if json_file_read_error != nil {
		fmt.Println(
			json_file_read_error)
	}

	fmt.Printf(
		"Successfully Opened %s\n", json_file_name)

	defer json_data.Close()

	json_data_byte_array, _ :=
		ioutil.ReadAll(
			json_data)

	return json_data_byte_array
}

func Pretty_print(json_data [][]interface{}) []byte {

	pretty_printed_json, _ := json.MarshalIndent(
		json_data,
		"",
		"	") //#TODO add pretty_printer to general utilities

	return pretty_printed_json

} //TODO - Stage 1 - move to json utilities
