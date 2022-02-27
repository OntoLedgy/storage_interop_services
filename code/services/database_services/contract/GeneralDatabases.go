package contract

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"

	"github.com/jmoiron/sqlx"
)

// GeneralDatabases represents a base "class" database_i_o_service - for all other concrete database_services
// it implements partly the Database interface
type GeneralDatabases struct {
	GetColumnsOfTableStmt *sqlx.Stmt
	*sqlx.DB
	*configurations.Settings
	Driver string
}

// Connect establishes a connection to the database_i_o_service with the given DSN.
// It pings the database_i_o_service to ensure it is reachable.
func (gdb *GeneralDatabases) Connect(dsn string) (err error) {
	gdb.DB, err = sqlx.Connect(gdb.Driver, dsn)
	if err != nil {
		usingPswd := "no"
		if gdb.Settings.Password != "" {
			usingPswd = "yes"
		}
		return fmt.Errorf(
			"could not connect to database_i_o_service (type=%q, user=%q, database_i_o_service=%q, host='%v:%v', using password: %v): %v",
			gdb.DbType, gdb.User, gdb.DbName, gdb.Host, gdb.Port, usingPswd, err,
		)
	}

	return gdb.Ping()
}

// Close closes the database_i_o_service connection
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

func (gdb *GeneralDatabases) BeginDatabaseTransaction() *sqlx.Tx {

	transaction := gdb.DB.MustBegin()

	return transaction
}
