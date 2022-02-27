package database_to_slice_interop_service

import (
	"database/sql"
	"fmt"
)

//----------------------------------------------

//#TODO - MOVE this into a ScanMsAccessRowset function in platform_specific_database_services pacakge

func Convert_rows_to_2d_slices(table_rowset *sql.Rows) [][]interface{} {

	table_columns, _ := table_rowset.Columns()

	var row_set_slice [][]interface{}
	var row_slice []interface{}

	rows := make([]interface{}, len(table_columns))

	for row_index, _ := range rows {
		var row_interface interface{}
		rows[row_index] = &row_interface
	}

	for table_rowset.Next() {

		table_rowset.Scan(rows...)

		for column_index, _ := range table_columns {
			var raw_value = *(rows[column_index].(*interface{}))
			row_slice = append(row_slice, raw_value)

		}
		row_set_slice = append(row_set_slice, row_slice)
		row_slice = nil

	}

	return row_set_slice
}

func Convert_rows_to_1d_slice(table_rowset *sql.Rows) []interface{} {

	table_columns, _ := table_rowset.Columns()

	var row_set_slice []interface{}
	var row_slice []string

	rows := make([]interface{}, len(table_columns))

	for row_index, _ := range rows {
		var row_interface interface{}
		rows[row_index] = &row_interface
	}

	for table_rowset.Next() {

		table_rowset.Scan(rows...)

		for column_index, _ := range table_columns {
			var raw_value = *(rows[column_index].(*interface{}))
			if raw_value != nil {
				row_slice = append(row_slice, raw_value.(string))
			} else {
				row_slice = nil
			}

		}
		if row_slice != nil {
			row_set_slice = append(row_set_slice, row_slice[0])
		}
		row_slice = nil

	}

	return row_set_slice
}

func Change_2d_interface_slice_to_string(two_d_interface [][]interface{}) [][]string {

	var return_string_slice [][]string
	var return_string_row []string

	for _, one_d_interface := range two_d_interface {

		for _, interface_element := range one_d_interface {

			switch interface_element.(type) {
			case int:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			case string:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			default:
				str := fmt.Sprintf("%v", interface_element)
				return_string_row = append(return_string_row, str)
			}

		}

		return_string_slice = append(return_string_slice, return_string_row)
		return_string_row = nil
	}

	return return_string_slice

}
