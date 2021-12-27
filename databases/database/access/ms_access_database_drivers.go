package access

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-adodb"
	"strings"
)

type MsAccessDatabaseDrivers struct {
	Database                *sqlx.DB
	OleDb12ConnectionString string
}

func (ms_access_driver *MsAccessDatabaseDrivers) Open() (*sqlx.DB, error) {

	//ms_access_driver := new(database.Driver)
	var database_open_error error

	connection_string := ms_access_driver.OleDb12ConnectionString

	connection_string_tokens := strings.Split(connection_string, ";")

	for _, connection_string_token := range connection_string_tokens {

		connection_string_variable_value_pair :=
			strings.
				Split(connection_string_token, "=")

		if len(
			connection_string_variable_value_pair) > 1 &&
			connection_string_variable_value_pair[0] == "Data Source" {

			database_encoding, database_encoding_error := ReadEncoding(connection_string_variable_value_pair[1])

			if database_encoding_error != nil {
				return nil, database_encoding_error
			}

			connection_string = connection_string + OleDb12ConnectionString_Suffix + database_encoding + `;`
		}
	}

	ms_access_driver.Database, database_open_error = sqlx.Open(database_driver_type, connection_string)

	if database_open_error != nil {
		return nil, database_open_error
	}

	database_open_error = ms_access_driver.Database.Ping()

	if database_open_error != nil {
		return nil, database_open_error
	}

	return ms_access_driver.Database, nil

}

func (ms_access_driver *MsAccessDatabaseDrivers) Close() error {

	ms_access_driver.Database.Close()

	return nil
}

func (ms_acess_driver *MsAccessDatabaseDrivers) Truncate_table(table_name string) {

	sql_string := truncate_table_sql_prefix + table_name

	ms_access_database_transaction := ms_acess_driver.Database.MustBegin()

	ms_access_database_transaction.MustExec(sql_string)

	ms_access_database_transaction.Commit()
}
