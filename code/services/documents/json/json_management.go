package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonToByteArray(
	jsonFileName string) []byte {

	jsonData, jsonFileReadError :=
		os.Open(
			jsonFileName)

	if jsonFileReadError != nil {
		fmt.Println(
			jsonFileReadError)
	}

	fmt.Printf(
		"Successfully Opened %s\n", jsonFileName)

	defer jsonData.Close()

	jsonDataByteArray, _ :=
		ioutil.ReadAll(
			jsonData)

	return jsonDataByteArray
}

func ReadJson(jsonFileName string) map[string]any {

	jsonDataByteArray := ReadJsonToByteArray(jsonFileName)

	var jsonDataStructure map[string]any

	err := json.Unmarshal(
		jsonDataByteArray,
		&jsonDataStructure)

	fmt.Println(err)

	return jsonDataStructure

}

func PrettyPrint(json_data [][]interface{}) []byte {

	prettyPrintedJson, _ := json.MarshalIndent(
		json_data,
		"",
		"	") //#TODO add pretty_printer to general utilities

	return prettyPrintedJson

} //TODO - Stage 1 - move to json utilities
