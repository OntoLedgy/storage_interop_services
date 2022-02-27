package postgresql

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/constants"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"

	// postgres database_i_o_service driver
	_ "github.com/lib/pq"
)

// PostgresqlDatabases implemenmts the Database interface with help of generalDatabase

// NewPostgresql creates a new PostgresqlDatabases database_i_o_service

func NewPostgresql(s *configurations.Settings) *PostgresqlDatabases {
	return &PostgresqlDatabases{
		GeneralDatabases: &contract.GeneralDatabases{
			Settings: s,
			Driver:   constants.DbTypeToDriverMap[s.DbType],
		},
		defaultUserName: "postgres",
	}
}
