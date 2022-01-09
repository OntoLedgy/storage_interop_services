package contract

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions/data_model"
)

//copied from https://medium.com/hackernoon/how-to-work-with-databases-in-golang-33b002aa8c47

type IDatabaseTransactions interface {
	Selectx(
		interface{},
		data_model.Queryx,
		...data_model.SelectOptions) error

	Get(
		interface{},
		data_model.Queryx) error

	Update(
		interface{}) error

	Delete(
		o interface{}) error

	Insert(
		o interface{}) error

	Count(
		data_model.Queryx) (
		int,
		error)
}
