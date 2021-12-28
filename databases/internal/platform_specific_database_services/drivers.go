package platform_specific_database_services

import "github.com/jmoiron/sqlx"

var drivers = make(
	map[string]Drivers)

type Drivers interface {
	Open() (*sqlx.DB, error)
	Close() error
}

func Register(driver *Drivers) (name string) {

	if driver == nil {
		panic("Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("Register called twice for driver " + name)
	}
	drivers[name] = driver
}
