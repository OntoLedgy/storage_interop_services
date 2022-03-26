package csv

import (
	"encoding/csv"
	"os"
)

func Write_2d_slice_set_to_csv_Obsolete(
	slice_to_write [][]string,
	csv_file *os.File) bool {

	for _, slice_row := range slice_to_write {

		Write1dSliceToCsv(
			slice_row,
			csv_file)

	}

	return true
}

func Write_1d_slice_to_csv_Obsolete(
	slice_to_write []string,
	csv_file *os.File) bool {

	writer_to_file :=
		csv.NewWriter(csv_file)

	writer_to_file.Write(
		slice_to_write)

	writer_to_file.Flush()

	return true

}
