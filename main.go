package main

import (
	"fmt"
	//"github.com/saintfish/chardet"
	"storage_interop_services/csv"
)

func main() {

	csv_file_name := `C:\OneDrive - BORO Engineering\BORO\Tools\Syntactic checker\Non-TEN_AllVersionsRenditions_20190104.csv`

	csv_file, csv_file_data := storage.Open_csv_file(csv_file_name)

	csv_dataset := storage.Read_csv_to_slice(csv_file, csv_file_data)

	/*for _, each := range csv_dataset {
		fmt.Println(each)
	}*/
	//fmt.Println(csv_dataset)

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
