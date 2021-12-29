package model

import (
	"database/sql/driver"
	database_contract "github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
	database_contract_implementation "github.com/OntoLedgy/storage_interop_services/code/services/databases/internal"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"time"
)

// copied from https://medium.com/hackernoon/how-to-work-with-databases-in-golang-33b002aa8c47

type Gender string

var (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (u *Gender) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b := value.([]byte)
	*u = Gender(b)
	return nil
}

func (u Gender) Value() (driver.Value, error) {
	return string(u), nil
}

type Person struct {
	PersonID     uuid.UUID `contract:"person_id"`
	FirstName    string    `contract:"first_name"`
	LastName     string    `contract:"last_name"`
	Active       bool      `contract:"active"`
	Gender       Gender    `contract:"gender"`
	ModifiedDate time.Time `contract:"modified_date"`
}

var (
	queryPersons      database_contract.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons"
	queryPersonByID   database_contract.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons WHERE person_id=:person_id"
	queryPersonInsert database_contract.Query = "INSERT INTO persons (person_id, first_name, last_name, gender, active, modified_date) VALUES (:person_id, :first_name, :last_name, :gender, :active, :modified_date)"
	queryPersonUpdate database_contract.Query = "UPDATE persons SET first_name=:first_name, last_name=:last_name, gender=:gender, modified_date=:modified_date, active=:active WHERE person_id=:person_id"
)

func QueryPersons(offset, count int) database_contract_implementation.Queryx {
	return database_contract_implementation.Queryx{
		Query:  queryPersons,
		Params: []interface{}{},
	}
}

func QueryPersonByID(personID utils.UUIDs) database_contract_implementation.Queryx {
	return database_contract_implementation.Queryx{
		Query: queryPersonByID,
		Params: []interface{}{
			personID,
		},
	}
}

func NewPerson() *Person {
	person_uuid, _ := utils.GetUUID(1, "")
	return &Person{PersonID: *person_uuid.UUID, ModifiedDate: time.Now()}
}

func (s *Person) Insert(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(queryPersonInsert), s)
	return err
}

func (s *Person) Update(tx *sqlx.Tx) error {
	s.ModifiedDate = time.Now()
	_, err := tx.NamedExec(string(queryPersonUpdate), s)
	return err
}

func (s *Person) Delete(tx *sqlx.Tx) error {
	s.Active = false
	return s.Update(tx)
}
