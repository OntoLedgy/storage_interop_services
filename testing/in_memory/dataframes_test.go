package in_memory

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
	"github.com/go-gota/gota/dataframe"
	"testing"
)

func TestDataframes(t *testing.T) {

	//csv_file := file_system_service.SelectFile()

	csv_data := csv.ReadCsvToSlice("D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\dataframe_test_data.csv", ",")

	testDataFrame := dataframes.DataFrames{
		DataFrame: dataframe.LoadRecords(
			csv_data),
	}

	testDataFrame.FillNa("TestNull")

	fmt.Println(testDataFrame)

}
