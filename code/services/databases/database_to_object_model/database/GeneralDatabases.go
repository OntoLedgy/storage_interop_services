package database

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
	"github.com/jmoiron/sqlx"
)

var (
	// dbTypeToDriverMap maps the database type to the driver names
	dbTypeToDriverMap = map[configurations.DatabaseType]string{
		configurations.DbTypePostgresql: "postgres",
		configurations.DbTypeMySQL:      "mysql",
		configurations.DbTypeSQLite:     "sqlite3",
	}
)

// GeneralDatabases represents a base "class" database - for all other concrete databases
// it implements partly the Database interface
type GeneralDatabases struct {
	GetColumnsOfTableStmt *sqlx.Stmt
	*sqlx.DB
	*configurations.Settings
	driver string
}

// New creates a new Database based on the given type in the settings.
func New(s *configurations.Settings) contract.Database {

	var db contract.Database

	switch s.DbType {
	case configurations.DbTypeSQLite:
		db = NewSQLite(s)
	case configurations.DbTypeMySQL:
		db = NewMySQL(s)
	case configurations.DbTypePostgresql:
		fallthrough
	default:
		db = NewPostgresql(s)
	}

	return db
}

// Connect establishes a connection to the database with the given DSN.
// It pings the database to ensure it is reachable.
func (gdb *GeneralDatabases) Connect(dsn string) (err error) {
	gdb.DB, err = sqlx.Connect(gdb.driver, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.Settings.Pswd != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf(
			"could not connect to database (type=%q, user=%q, database=%q, host='%v:%v', using password: %v): %v",
			gdb.DbType, gdb.User, gdb.DbName, gdb.Host, gdb.Port, usingPswd, err,
		)
	}

	return gdb.Ping()
}

// Close closes the database connection
func (gdb *GeneralDatabases) Close() error {
	return gdb.DB.Close()
}

// IsNullable returns true if column is a nullable one
func (gdb *GeneralDatabases) IsNullable(column object_model.Column) bool {
	return column.IsNullable == "YES"
}

// IsStringInSlice checks if needle (string) is in haystack ([]string)
func (gdb *GeneralDatabases) IsStringInSlice(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
