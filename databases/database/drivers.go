package database

import "github.com/jmoiron/sqlx"

var drivers = make(map[string]Driver)

type Driver interface {
	Open() (*sqlx.DB, error)
	Close() error
}

func Register(name string, driver Driver) {

	if driver == nil {
		panic("Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("Register called twice for driver " + name)
	}
	drivers[name] = driver
}
