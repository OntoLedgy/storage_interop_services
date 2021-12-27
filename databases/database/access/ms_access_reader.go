package access

import (
	"bytes"
	"database/sql"
	"fmt"
)

func ReadMsAccessTable(ms_access_database *MsAccessDatabaseDrivers, table_name string) *sql.Rows {

	var cmd bytes.Buffer

	sql_string := select_all_sql_prefix + table_name

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Database.Query(cmd.String())
	if query_error != nil {
		panic(query_error)
	}

	columns, _ := rowset.Columns()

	for _, column := range columns {

		fmt.Println(column)
	}

	rows := make([]interface{}, len(columns))

	for row_index, _ := range rows {
		var row_interface interface{}
		rows[row_index] = &row_interface
	}

	//TODO - return full dataset as a slice of [][]interface{}

	return rowset
}

func ReadMsAccessColumns(ms_access_database *MsAccessDatabaseDrivers, table_name string, column_list []string) *sql.Rows {

	var cmd bytes.Buffer

	var sql_column_names string

	for _, column_name := range column_list {
		sql_column_names += "[" + column_name + "],"
	}

	sql_string := select_sql_prefix + sql_column_names[0:len(sql_column_names)-1] + select_sql_from + "[" + table_name + "]"

	fmt.Println(sql_string)

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Database.Query(cmd.String())

	if query_error != nil {
		panic(query_error)
		ms_access_database.Database.DB.Close()
	}

	//TODO - return full dataset as a slice of [][]interface{}

	return rowset
}

func ReadMsAccessColumn(ms_access_database *MsAccessDatabaseDrivers, table_name string, column_name string) *sql.Rows {

	var cmd bytes.Buffer

	var sql_column_names string

	/*for _, column_names := range column_list {
		sql_column_names += column_names + ","
	}*/

	sql_string := select_sql_prefix + sql_column_names + select_sql_from + table_name

	fmt.Println(sql_string)

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Database.Query(cmd.String())

	if query_error != nil {
		panic(query_error)
	}

	//TODO - return full dataset as a slice of [][]interface{}

	return rowset
}
