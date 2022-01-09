package contract

import (
	"github.com/jmoiron/sqlx"
)

type IDatabases interface {
	ConnectDatabase()
	BeginDatabaseTransaction() *sqlx.Tx
	CloseDatabase()
}
