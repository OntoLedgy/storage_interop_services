package internal

import (
	"errors"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/databases/contract"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"reflect"
)

type DatabaseInteropServices struct {
	//TODO : add configuraiton strucure here
	*sqlx.Tx
}

func (Database_interop_service *DatabaseInteropServices) Update() {

}

//TODO : replace with standard logging services
var log = logging.MustGetLogger("contract")

var (
	ErrNoGetterFound   = errors.New("No getter found")
	ErrNoDeleterFound  = errors.New("No deleter found")
	ErrNoSelecterFound = errors.New("No getter found")
	ErrNoUpdaterFound  = errors.New("No updater found")
	ErrNoInserterFound = errors.New("No inserter found")
)

//TODO : encapsulate into an interface

type Queryx struct {
	Query  contract.Query
	Params []interface{}
}

type Databases struct {
	*sqlx.DB
}

type DatabaseTransactions struct {
	*sqlx.Tx
}

type limitOption struct {
	offset int
	count  int
}

func Limit(
	offset, count int) contract.SelectOptions {
	return &limitOption{offset, count}
}

func (o *limitOption) Wrap(
	query string,
	params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("SELECT a.* FROM (%s) a LIMIT ?, ?", query)
	params = append(params, o.offset)
	params = append(params, o.count)
	return query, params
}

func (database_transaction *DatabaseTransactions) Selectx(
	o interface{},
	qx Queryx,
	options ...contract.SelectOptions) error {
	q := string(qx.Query)
	params := qx.Params
	log.Debug(q)
	for _, option := range options {
		q, params = option.Wrap(q, params)
	}
	if u, ok := o.(contract.Selecter); ok {
		return u.Select(database_transaction.Tx, contract.Query(q), params...)
	}
	stmt, err := database_transaction.Preparex(q)
	if err != nil {
		return err
	}
	return stmt.Select(o, params...)
}

func (database_transaction *DatabaseTransactions) Getx(
	o interface{},
	qx Queryx) error {
	if u, ok := o.(contract.Getter); ok {
		return u.Get(database_transaction.Tx, qx.Query, qx.Params...)
	}
	stmt, err := database_transaction.Preparex(string(qx.Query))
	if err != nil {
		return err
	}
	return stmt.Get(o, qx.Params...)
}

func (database_transaction *DatabaseTransactions) Get(
	o interface{},
	query contract.Query,
	params ...interface{}) error {
	if u, ok := o.(contract.Getter); ok {
		return u.Get(database_transaction.Tx, query, params...)
	}
	stmt, err := database_transaction.Preparex(string(query))
	if err != nil {
		return err
	}
	return stmt.Get(o, params...)
}

func (database_transaction *DatabaseTransactions) Update(
	o interface{}) error {

	if u, ok := o.(contract.Updater); ok {
		return u.Update(database_transaction.Tx)
	}

	log.Debug(
		"No updater found for object: %s",
		reflect.TypeOf(o))

	return ErrNoUpdaterFound
}

func (database_transaction *DatabaseTransactions) Delete(
	o interface{}) error {
	if u, ok := o.(contract.Deleter); ok {
		return u.Delete(database_transaction.Tx)
	}
	log.Debug("No deleter found for object: %s", reflect.TypeOf(o))
	return ErrNoDeleterFound
}

func (database_transaction *DatabaseTransactions) Insert(
	o interface{}) error {
	if u, ok := o.(contract.Inserter); ok {
		err := u.Insert(database_transaction.Tx)
		if err != nil {
			log.Error(err.Error())
		}
		return err
	}
	log.Debug("No inserter found for object: %s", reflect.TypeOf(o))
	return ErrNoInserterFound
}

func (db *Databases) Begin() *DatabaseTransactions {
	tx := db.MustBegin()
	return &DatabaseTransactions{tx}
}

func (database_transaction *DatabaseTransactions) Countx(
	qx Queryx) (int, error) {
	stmt, err := database_transaction.Preparex(fmt.Sprintf("SELECT COUNT(*) FROM (%s) q", string(qx.Query)))
	if err != nil {
		return 0, err
	}
	count := 0
	err = stmt.Get(&count, qx.Params...)
	return count, err
}
