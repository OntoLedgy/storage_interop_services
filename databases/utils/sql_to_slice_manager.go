package utils

import "database/sql"

//----------------------------------------------

//#TODO - MOVE this into a ScanMsAccessRowset function in database pacakge

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
