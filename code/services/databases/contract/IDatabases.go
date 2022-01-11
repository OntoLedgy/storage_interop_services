package contract

import (
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/jmoiron/sqlx"
)

type IDatabases interface {
	ConnectDatabase()
	BeginDatabaseTransaction() *sqlx.Tx
	CloseDatabase()
}

// Database interface for the concrete databases
type Database interface {
	DSN() string
	Connect() (err error)
	Close() (err error)
	GetDriverImportLibrary() string

	GetTables() (tables []*object_model.Table, err error)
	PrepareGetColumnsOfTableStmt() (err error)
	GetColumnsOfTable(table *object_model.Table) (err error)

	IsPrimaryKey(column object_model.Column) bool
	IsAutoIncrement(column object_model.Column) bool
	IsNullable(column object_model.Column) bool

	GetStringDatatypes() []string
	IsString(column object_model.Column) bool

	GetTextDatatypes() []string
	IsText(column object_model.Column) bool

	GetIntegerDatatypes() []string
	IsInteger(column object_model.Column) bool

	GetFloatDatatypes() []string
	IsFloat(column object_model.Column) bool

	GetTemporalDatatypes() []string
	IsTemporal(column object_model.Column) bool
	GetTemporalDriverDataType() string

	// TODO pg: bitstrings, enum, range, other special types
	// TODO mysql: bit, enums, set
}
