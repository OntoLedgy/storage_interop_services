package testing

import (
	"fmt"
	files "github.com/OntoLedgy/storage_interop_services/files"
	"testing"
)

func TestStandard(t *testing.T) {

	fmt.Println("testing file")

	files.Create_file_if_does_not_exist_else_delete_it("C:\\S\\test.text")

}
