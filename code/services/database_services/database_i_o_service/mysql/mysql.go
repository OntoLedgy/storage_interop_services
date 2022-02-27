package mysql

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/constants"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	object_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"strings"

	// MySQL database_i_o_service Driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQL implemenmts the Database interface with help of generalDatabase
type MySQL struct {
	*contract.GeneralDatabases
}

// NewMySQL creates a new MySQL database_i_o_service
func NewMySQL(s *configurations.Settings) *MySQL {
	return &MySQL{
		GeneralDatabases: &contract.GeneralDatabases{
			Settings: s,
			Driver:   constants.DbTypeToDriverMap[s.DbType],
		},
	}
}

// Connect connects to the database_i_o_service by the given data source name (dsn) of the concrete database_i_o_service
func (mysql *MySQL) Connect() error {
	return mysql.GeneralDatabases.Connect(mysql.DSN())
}

// DSN creates the DSN String to connect to this database_i_o_service
func (mysql *MySQL) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		mysql.Settings.User, mysql.Settings.Password, mysql.Settings.Host, mysql.Settings.Port, mysql.Settings.DbName)
}

// GetDriverImportLibrary returns the golang sql Driver specific fot the MySQL database_i_o_service
func (mysql *MySQL) GetDriverImportLibrary() string {
	return `"github.com/go-sql-Driver/mysql"`
}

// GetTables gets all tables for a given database_i_o_service by name
func (mysql *MySQL) GetTables() (tables []*object_model2.Table, err error) {

	err = mysql.Select(&tables, `
		SELECT table_name AS table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		AND table_schema = ?
		ORDER BY table_name
	`, mysql.DbName)

	if mysql.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> schema: %q\r\n", mysql.DbName)
		}
	}

	return tables, err
}

// PrepareGetColumnsOfTableStmt prepares the statement for retrieving the columns of a specific table for a given database_i_o_service
func (mysql *MySQL) PrepareGetColumnsOfTableStmt() (err error) {

	mysql.GetColumnsOfTableStmt, err = mysql.Preparex(`
		SELECT
		  ordinal_position AS ordinal_position,
		  column_name AS column_name,
		  data_type AS data_type,
		  column_default AS column_default,
		  is_nullable AS is_nullable,
		  character_maximum_length AS character_maximum_length,
		  numeric_precision AS numeric_precision,
		  column_key AS column_key,
		  extra AS extra
		FROM information_schema.columns
		WHERE table_name = ?
		AND table_schema = ?
		ORDER BY ordinal_position
	`)

	return err
}

// GetColumnsOfTable executes the statement for retrieving the columns of a specific table for a given database_i_o_service
func (mysql *MySQL) GetColumnsOfTable(table *object_model2.Table) (err error) {

	err = mysql.GetColumnsOfTableStmt.Select(&table.Columns, table.Name, mysql.DbName)

	if mysql.Settings.Verbose {
		if err != nil {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> schema: %q\r\n", mysql.Schema)
			fmt.Printf("> dbName: %q\r\n", mysql.DbName)
		}
	}

	return err
}

// IsPrimaryKey checks if column belongs to primary key
func (mysql *MySQL) IsPrimaryKey(column object_model2.Column) bool {
	return strings.Contains(column.ColumnKey, "PRI")
}

// IsAutoIncrement checks if column is a auto_increment column
func (mysql *MySQL) IsAutoIncrement(column object_model2.Column) bool {
	return strings.Contains(column.Extra, "auto_increment")
}

// GetStringDatatypes returns the string datatypes for the MySQL database_i_o_service
func (mysql *MySQL) GetStringDatatypes() []string {
	return []string{
		"char",
		"varchar",
		"binary",
		"varbinary",
	}
}

// IsString returns true if colum is of type string for the MySQL database_i_o_service
func (mysql *MySQL) IsString(column object_model2.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetStringDatatypes())
}

// GetTextDatatypes returns the text datatypes for the MySQL database_i_o_service
func (mysql *MySQL) GetTextDatatypes() []string {
	return []string{
		"text",
		"blob",
	}
}

// IsText returns true if colum is of type text for the MySQL database_i_o_service
func (mysql *MySQL) IsText(column object_model2.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetTextDatatypes())
}

// GetIntegerDatatypes returns the integer datatypes for the MySQL database_i_o_service
func (mysql *MySQL) GetIntegerDatatypes() []string {
	return []string{
		"tinyint",
		"smallint",
		"mediumint",
		"int",
		"bigint",
	}
}

// IsInteger returns true if colum is of type integer for the MySQL database_i_o_service
func (mysql *MySQL) IsInteger(column object_model2.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetIntegerDatatypes())
}

// GetFloatDatatypes returns the float datatypes for the MySQL database_i_o_service
func (mysql *MySQL) GetFloatDatatypes() []string {
	return []string{
		"numeric",
		"decimal",
		"float",
		"real",
		"double precision",
	}
}

// IsFloat returns true if colum is of type float for the MySQL database_i_o_service
func (mysql *MySQL) IsFloat(column object_model2.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetFloatDatatypes())
}

// GetTemporalDatatypes returns the temporal datatypes for the MySQL database_i_o_service
func (mysql *MySQL) GetTemporalDatatypes() []string {
	return []string{
		"time",
		"timestamp",
		"date",
		"datetime",
		"year",
	}
}

// IsTemporal returns true if colum is of type temporal for the MySQL database_i_o_service
func (mysql *MySQL) IsTemporal(column object_model2.Column) bool {
	return mysql.IsStringInSlice(column.DataType, mysql.GetTemporalDatatypes())
}

// GetTemporalDriverDataType returns the time data type specific for the MySQL database_i_o_service
func (mysql *MySQL) GetTemporalDriverDataType() string {
	return "mysql.NullTime"
}
