package model

import (
	"database/sql/driver"
	"database_manager/general_database"
	"database_manager/utils"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"time"
)

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
	PersonID     uuid.UUID `general_database:"person_id"`
	FirstName    string    `general_database:"first_name"`
	LastName     string    `general_database:"last_name"`
	Active       bool      `general_database:"active"`
	Gender       Gender    `general_database:"gender"`
	ModifiedDate time.Time `general_database:"modified_date"`
}

var (
	queryPersons      db.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons"
	queryPersonByID   db.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons WHERE person_id=:person_id"
	queryPersonInsert db.Query = "INSERT INTO persons (person_id, first_name, last_name, gender, active, modified_date) VALUES (:person_id, :first_name, :last_name, :gender, :active, :modified_date)"
	queryPersonUpdate db.Query = "UPDATE persons SET first_name=:first_name, last_name=:last_name, gender=:gender, modified_date=:modified_date, active=:active WHERE person_id=:person_id"
)

func QueryPersons(offset, count int) db.Queryx {
	return db.Queryx{
		Query:  queryPersons,
		Params: []interface{}{},
	}
}
func QueryPersonByID(personID utils.UUIDs) db.Queryx {
	return db.Queryx{
		Query: queryPersonByID,
		Params: []interface{}{
			personID,
		},
	}
}
func NewPerson() *Person {
	person_uuid, _ := utils.GetUUID(1, "")
	return &Person{PersonID: person_uuid, ModifiedDate: time.Now()}
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
