package contract

import (
	object_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/jmoiron/sqlx"
)

//type IDatabases interface {
//	ConnectDatabase()
//	BeginDatabaseTransaction() *sqlx.Tx
//	CloseDatabase()
//}

// Database interface for the concrete database_services
type IDatabases interface {
	DSN() string
	Connect() (err error)
	Close() (err error)
	GetDriverImportLibrary() string

	GetTables() (tables []*object_model2.Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *object_model2.Table) (err error)

	IsPrimaryKey(column object_model2.Column) bool
	IsAutoIncrement(column object_model2.Column) bool
	IsNullable(column object_model2.Column) bool

	GetStringDatatypes() []string
	IsString(column object_model2.Column) bool

	GetTextDatatypes() []string
	IsText(column object_model2.Column) bool

	GetIntegerDatatypes() []string
	IsInteger(column object_model2.Column) bool

	GetFloatDatatypes() []string
	IsFloat(column object_model2.Column) bool

	GetTemporalDatatypes() []string
	IsTemporal(column object_model2.Column) bool
	GetTemporalDriverDataType() string

	// TODO pg: bitstrings, enum, range, other special types
	// TODO mysql: bit, enums, set

	BeginDatabaseTransaction() *sqlx.Tx
}
