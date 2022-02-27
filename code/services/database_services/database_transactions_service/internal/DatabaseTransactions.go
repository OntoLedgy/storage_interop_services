package internal

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_transactions_service/object_model"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type DatabaseTransactions struct {
	*sqlx.Tx
}

func (database_transaction *DatabaseTransactions) Selectx(
	transactionObject interface{},
	query object_model.Queryx,
	options ...object_model.SelectOptions) error {

	queryString :=
		string(query.Query)

	queryParameters :=
		query.Params

	fmt.Println(
		queryString)

	for _, option := range options {
		queryString, queryParameters =
			option.Wrap(
				queryString,
				queryParameters)
	}

	var queryResults = database_transaction.Select(
		transactionObject,
		queryString,
		queryParameters...)

	return queryResults

}

func (database_transaction *DatabaseTransactions) Get(
	o interface{},
	qx object_model.Queryx) error {

	if u, ok := o.(Getter); ok {
		return u.Get(
			database_transaction.Tx,
			qx.Query,
			qx.Params...)
	}
	stmt, err := database_transaction.Preparex(string(qx.Query))
	if err != nil {
		return err
	}
	return stmt.Get(o, qx.Params...)
}

func (database_transaction *DatabaseTransactions) Update(
	o interface{}) error {

	if u, ok := o.(Updater); ok {
		return u.Update(database_transaction.Tx)
	}

	fmt.Printf(
		"No updater found for object: %s",
		reflect.TypeOf(o))

	return ErrNoUpdaterFound
}

func (database_transaction *DatabaseTransactions) Delete(
	o interface{}) error {
	if u, ok := o.(Deleter); ok {
		return u.Delete(database_transaction.Tx)
	}
	fmt.Printf(
		"No deleter found for object: %s",
		reflect.TypeOf(o))

	return ErrNoDeleterFound
}

func (database_transaction *DatabaseTransactions) Insert(
	o interface{}) error {

	if u, ok := o.(Inserter); ok {
		err := u.Insert(database_transaction.Tx)
		if err != nil {
			fmt.Println(err.Error())
		}
		return err
	}
	fmt.Printf(
		"No inserter found for object: %s", reflect.TypeOf(o))

	return ErrNoInserterFound
}

func (database_transaction *DatabaseTransactions) Count(
	qx object_model.Queryx) (
	int,
	error) {
	stmt, err := database_transaction.Preparex(fmt.Sprintf("SELECT COUNT(*) FROM (%s) q", string(qx.Query)))
	if err != nil {
		return 0, err
	}
	count := 0
	err = stmt.Get(&count, qx.Params...)
	return count, err
}
