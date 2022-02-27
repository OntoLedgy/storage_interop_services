package contract

import (
	data_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/object_model"
)

//copied from https://medium.com/hackernoon/how-to-work-with-databases-in-golang-33b002aa8c47

type IDatabaseTransactions interface {
	Selectx(
		interface{},
		data_model2.Queryx,
		...data_model2.SelectOptions) error

	Get(
		interface{},
		data_model2.Queryx) error

	Update(
		interface{}) error

	Delete(
		o interface{}) error

	Insert(
		o interface{}) error

	Count(
		data_model2.Queryx) (
		int,
		error)
}
