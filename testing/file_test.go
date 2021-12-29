package testing

import (
	"fmt"
	files "github.com/OntoLedgy/storage_interop_services/code/services/disk/files"
	storage "github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"testing"
)

func TestStandard(t *testing.T) {

	fmt.Println("testing file")

	files.Create_file_if_does_not_exist_else_delete_it("C:\\S\\test.text")

	//TODO move this to testing

	csv_file_name :=
		`C:\S\test.text`

	csv_file, csv_file_data := storage.Open_csv_file(csv_file_name)

	csv_dataset :=
		storage.Read_csv_to_slice(
			csv_file,
			csv_file_data,
			",")

	for row_index, row := range csv_dataset {
		for col_index, col := range row {
			for string_index, col_chars := range col {
				if row_index == 2 {
					fmt.Printf("row_index:%v col index :%v : col string:[%s] string_index:%v char value: %v, char:%c\n ", row_index, col_index, col, string_index, col_chars, col_chars)
				}
			}

			//fmt.Printf("%v:, %30q\n ", col_index, col)

		}
		//storage_interop_services.Write_slice_to_csv_split_by_column(4, csv_file)
	}

}
