package constants

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
)

var (
	// DbTypeToDriverMap maps the database_i_o_service type to the Driver names
	DbTypeToDriverMap = map[configurations.DatabaseType]string{
		configurations.DbTypePostgresql: "postgres",
		configurations.DbTypeMySQL:      "mysql",
		configurations.DbTypeSQLite:     "sqlite3",
		configurations.DbTypeAccess:     "access",
	}
)
