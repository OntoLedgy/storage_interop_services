package database_transactions

import (
	database_transaction_contract "github.com/OntoLedgy/storage_interop_services/code/services/database_transactions/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions/internal"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
)

type DatabaseTransactionFactory struct {
	contract.IDatabases
}

func (databaseTransactionFactory *DatabaseTransactionFactory) Create() database_transaction_contract.IDatabaseTransactions {

	databaseTransaction := new(internal.DatabaseTransactions)

	databaseTransaction.Tx = databaseTransactionFactory.BeginDatabaseTransaction()

	return databaseTransaction

}
