package internal

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/object_model"
	"github.com/jmoiron/sqlx"
)

type Updater interface {
	Update(
		*sqlx.Tx) error
}

type Inserter interface {
	Insert(
		*sqlx.Tx) error
}

type Selecter interface {
	Select(
		*sqlx.Tx,
		object_model.Queries,
		...interface{}) error
}

type Deleter interface {
	Delete(*sqlx.Tx) error
}

type Getter interface {
	Get(
		*sqlx.Tx,
		object_model.Queries,
		...interface{}) error
}
