package database_transactions_service

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	database_transaction_contract "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/internal"
)

type DatabaseTransactionFactory struct {
	contract.IDatabases
}

func (databaseTransactionFactory *DatabaseTransactionFactory) Create() database_transaction_contract.IDatabaseTransactions {

	databaseTransaction := new(internal.DatabaseTransactions)

	databaseTransaction.Tx = databaseTransactionFactory.BeginDatabaseTransaction()

	return databaseTransaction

}
