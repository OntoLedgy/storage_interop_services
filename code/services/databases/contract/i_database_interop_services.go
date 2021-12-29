package contract

import (
	"github.com/jmoiron/sqlx"
)

//copied from https://medium.com/hackernoon/how-to-work-with-databases-in-golang-33b002aa8c47

type IDatabaseInteropServices interface {
	//TODO : add functions here

	Updater
	Inserter
	Selecter
	Getter
	Deleter
	SelectOptions
}

type Query string

type Updater interface {
	Update(*sqlx.Tx) error
}

type Inserter interface {
	Insert(*sqlx.Tx) error
}

type Selecter interface {
	Select(*sqlx.Tx, Query, ...interface{}) error
}

type Getter interface {
	Get(*sqlx.Tx, Query, ...interface{}) error
}

type Deleter interface {
	Delete(*sqlx.Tx) error
}

type SelectOptions interface {
	Wrap(string, []interface{}) (string, []interface{})
}
