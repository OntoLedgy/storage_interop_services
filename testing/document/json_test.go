package document

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/json"
	"testing"
)

func TestJsonIO(t *testing.T) {
	filePathAndName :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\sample.json"

	byteArray := json.ReadJsonToByteArray(
		filePathAndName)

	jsonData := json.ReadJson(
		filePathAndName)

	println(byteArray)

	println(jsonData)

}
