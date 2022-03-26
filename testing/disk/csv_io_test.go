package disk

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/file_system_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"testing"
)

func TestCsvIO(t *testing.T) {

	fileNameAndPath :=
		file_system_service.CreateFile()

	csv_file, csv_file_data := csv.OpenCsvFile(fileNameAndPath.AbsolutePathString())

	csv_dataset :=
		csv.Read_csv_to_slice(
			csv_file,
			csv_file_data,
			",")

	for row_index, row := range csv_dataset {
		for col_index, col := range row {
			fmt.Printf("%v:, %30q\n ", col_index, col)
			for string_index, col_chars := range col {
				fmt.Printf("row_index:%v col index :%v : col string:[%s] string_index:%v char value: %v, char:%c\n ",
					row_index,
					col_index,
					col,
					string_index,
					col_chars,
					col_chars)

			}
		}
	}

}
