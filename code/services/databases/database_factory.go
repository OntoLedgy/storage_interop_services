package databases

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/internal/postgresql"
)

type DatabaseFactory struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	DriverName   string
}

func (databaseFactory *DatabaseFactory) Create() contract.IDatabases {

	if databaseFactory.DriverName == "postgres" {

		database := new(postgresql.PostgresDatabases)
		database.Host = databaseFactory.Host
		database.Port = databaseFactory.Port
		database.User = databaseFactory.User
		database.Password = databaseFactory.Password
		database.DatabaseName = databaseFactory.DatabaseName
		database.DriverName = databaseFactory.DriverName

		return database
	}

	return nil

}
