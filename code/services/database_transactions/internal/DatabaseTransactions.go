package internal

import (
	"errors"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions/data_model"

	"github.com/jmoiron/sqlx"
	"reflect"
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
		data_model.Queries,
		...interface{}) error
}

type Deleter interface {
	Delete(*sqlx.Tx) error
}

type Getter interface {
	Get(
		*sqlx.Tx,
		data_model.Queries,
		...interface{}) error
}

type DatabaseTransactions struct {
	*sqlx.Tx
}

var (
	ErrNoGetterFound   = errors.New("No getter found")
	ErrNoDeleterFound  = errors.New("No deleter found")
	ErrNoSelecterFound = errors.New("No getter found")
	ErrNoUpdaterFound  = errors.New("No updater found")
	ErrNoInserterFound = errors.New("No inserter found")
)

func (database_transaction *DatabaseTransactions) Selectx(
	transactionObject interface{},
	query data_model.Queryx,
	options ...data_model.SelectOptions) error {

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
	qx data_model.Queryx) error {

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
	qx data_model.Queryx) (
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
