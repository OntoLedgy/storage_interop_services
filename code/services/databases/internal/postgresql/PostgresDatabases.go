package postgresql

import (
	"fmt"
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//TODO - try out https://github.com/jackc/pgx
//TODO - review comments tips from https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4

type PostgresDatabases struct {
	*sqlx.DB
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	DriverName   string
}

func (database *PostgresDatabases) ConnectDatabase() {
	// connection string
	psqlconn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		database.User,
		database.Password,
		database.Host,
		//database.Port,
		database.DatabaseName)

	// open database

	databaseHandle :=
		sqlx.MustConnect(
			database.DriverName,
			psqlconn)

	fmt.Printf("Connected to host:%s, port: %v database: %s using driver:%s\n",
		database.Host,
		database.Port,
		database.DatabaseName,
		database.DriverName)

	database.DB = databaseHandle

}

func (database *PostgresDatabases) BeginDatabaseTransaction() *sqlx.Tx {

	transaction := database.MustBegin()

	return transaction
}

func (database *PostgresDatabases) CloseDatabase() {
	// close database
	defer database.Close()

	// check database
	var err = database.Ping()

	if err != nil {
		log.Error("error")
	}

}
