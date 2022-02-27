package access

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/access/internal"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MsAccessDatabase struct {
	*contract.GeneralDatabases
	*sqlx.DB
	Host                   string
	Port                   int
	User                   string
	Password               string
	DriverName             string
	SystemDatabaseFileName string
	MsAccessDriver         MsAccessDatabaseDrivers
}

//func (database *MsAccessDatabase) ConnectDatabase_Deprecate() {
//	// connection string
//	//MsAccessDriver := new(MsAccessDatabaseDrivers)
//
//	var databaseOpenError error
//
//	database.MsAccessDriver.OleDb12ConnectionString =
//		internal.OleDb12ConnectionString_Prefix +
//			database.GeneralDatabases.Settings.DbName +
//			internal.SystemDatabaseReference +
//			database.SystemDatabaseFileName
//
//	// open database_i_o_service
//	database.DB, databaseOpenError =
//		database.MsAccessDriver.Open()
//
//	if databaseOpenError != nil {
//		panic(databaseOpenError)
//	}
//
//	fmt.Printf("Connected database_i_o_service: %s using driver:%s\n",
//		database.GeneralDatabases.Settings.DbName,
//		database.DriverName)
//
//}

func (database *MsAccessDatabase) BeginDatabaseTransaction() *sqlx.Tx {

	transaction := database.MustBegin()

	return transaction
}

func (database *MsAccessDatabase) DSN() string {
	return "pass"
}
func (database *MsAccessDatabase) Connect() error {

	// connection string
	//MsAccessDriver := new(MsAccessDatabaseDrivers)

	var databaseOpenError error

	database.MsAccessDriver.OleDb12ConnectionString =
		internal.OleDb12ConnectionString_Prefix +
			database.GeneralDatabases.Settings.DbName +
			internal.SystemDatabaseReference +
			database.SystemDatabaseFileName + ";"

	// open database_i_o_service
	database.DB, databaseOpenError =
		database.MsAccessDriver.Open()

	if databaseOpenError != nil {
		panic(databaseOpenError)
	}

	fmt.Printf("Connected database_i_o_service: %s using driver:%s\n",
		database.GeneralDatabases.Settings.DbName,
		database.DriverName)

	return databaseOpenError
}
func (database *MsAccessDatabase) Close() error {

	// close database_i_o_service
	err := database.DB.Close()

	return err
}
func (database *MsAccessDatabase) GetDriverImportLibrary() string {
	return "pass"
}

func (database *MsAccessDatabase) GetTables() (tables []*object_model.Table, err error) {

	err = database.Select(
		&tables,
		`SELECT Name AS table_name FROM MSysObjects WHERE Type = 1 AND Flags = 0`)

	if database.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", database.Schema)
		}
	}

	return tables, err

}

func (database *MsAccessDatabase) PrepareGetColumnsOfTableStmt() (err error) {
	return nil
}

func (database *MsAccessDatabase) GetColumnsOfTable(table *object_model.Table) (err error) {
	return nil
}

func (database *MsAccessDatabase) IsPrimaryKey(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) IsAutoIncrement(column object_model.Column) bool {
	return false
}
func (database *MsAccessDatabase) IsNullable(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) GetStringDatatypes() []string {
	return []string{}
}
func (database *MsAccessDatabase) IsString(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) GetTextDatatypes() []string {
	return []string{}
}
func (database *MsAccessDatabase) IsText(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) GetIntegerDatatypes() []string {
	return nil
}
func (database *MsAccessDatabase) IsInteger(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) GetFloatDatatypes() []string {
	return nil
}
func (database *MsAccessDatabase) IsFloat(column object_model.Column) bool {
	return false
}

func (database *MsAccessDatabase) GetTemporalDatatypes() []string {
	return []string{}
}
func (database *MsAccessDatabase) IsTemporal(column object_model.Column) bool {
	return false
}
func (database *MsAccessDatabase) GetTemporalDriverDataType() string {
	return "pass"
}

//func (database *MsAccessDatabase) CloseDatabase() {
//	// close database_i_o_service
//	defer database.Close()
//
//	// check database_i_o_service
//	var err = database.Ping()
//
//	if err != nil {
//		log.Error("error")
//	}
//}

//TODO - move to Transaction Level

func (database *MsAccessDatabase) TruncateTable(table_name string) {

	sql_string := internal.TruncateTableSqlPrefix + table_name

	ms_access_database_transaction := database.MustBegin()

	ms_access_database_transaction.MustExec(sql_string)

	ms_access_database_transaction.Commit()
}

func ReadMsAccessTable(ms_access_database *MsAccessDatabase, table_name string) *sql.Rows {

	var cmd bytes.Buffer

	sql_string := internal.SelectAllSqlPrefix + table_name

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Query(cmd.String())
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

func ReadMsAccessColumns(ms_access_database *MsAccessDatabase, table_name string, column_list []string) *sql.Rows {

	var cmd bytes.Buffer

	var sql_column_names string

	for _, column_name := range column_list {
		sql_column_names += "[" + column_name + "],"
	}

	sql_string := internal.SelectAllSqlPrefix + sql_column_names[0:len(sql_column_names)-1] + internal.SelectAllSqlPrefix + "[" + table_name + "]"

	fmt.Println(sql_string)

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Query(cmd.String())

	if query_error != nil {
		panic(query_error)
		ms_access_database.Close()
	}

	//TODO - return full dataset as a slice of [][]interface{}

	return rowset
}

func ReadMsAccessColumn(ms_access_database *MsAccessDatabase, table_name string, column_name string) *sql.Rows {

	var cmd bytes.Buffer

	var sql_column_names string

	/*for _, column_names := range column_list {
		sql_column_names += column_names + ","
	}*/

	sql_string := internal.SelectAllSqlPrefix + sql_column_names + internal.SelectAllSqlPrefix + table_name

	fmt.Println(sql_string)

	cmd.WriteString(sql_string)

	rowset, query_error := ms_access_database.Query(cmd.String())

	if query_error != nil {
		panic(query_error)
	}

	//TODO - return full dataset as a slice of [][]interface{}

	return rowset
}
