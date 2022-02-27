package database_i_o_service

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/access"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/mysql"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/postgresql"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/sqlite"
)

type DatabaseFactory struct {
	Host           string
	Port           int
	User           string
	Password       string
	DatabaseName   string
	DriverName     string
	DatabaseType   configurations.DatabaseType
	SystemDatabase string
}

func (databaseFactory *DatabaseFactory) New() contract.IDatabases {
	settingsFactory := &configurations.SettingsFactory{}

	settings := settingsFactory.Create()

	settings.Host = databaseFactory.Host
	settings.Port = databaseFactory.Port
	settings.User = databaseFactory.User
	settings.Password = databaseFactory.Password
	settings.DbName = databaseFactory.DatabaseName
	settings.DbType = databaseFactory.DatabaseType

	var db contract.IDatabases

	switch settings.DbType {
	case configurations.DbTypeSQLite:
		db = sqlite.NewSQLite(settings)
	case configurations.DbTypeMySQL:
		db = mysql.NewMySQL(settings)
	case configurations.DbTypePostgresql:
		db = postgresql.NewPostgresql(settings)
	case configurations.DbTypeAccess:
		db = access.NewMsAccessDatabase(settings, databaseFactory.SystemDatabase)
	default:
		return nil
	}

	return db

}

// New creates a new Database based on the given type in the settings.
func New(s *configurations.Settings) contract.IDatabases {

	var db contract.IDatabases

	switch s.DbType {
	case configurations.DbTypeSQLite:
		db = sqlite.NewSQLite(s)
	case configurations.DbTypeMySQL:
		db = mysql.NewMySQL(s)
	case configurations.DbTypePostgresql:
		db = postgresql.NewPostgresql(s)
	default:
		return nil
	}

	return db
}
