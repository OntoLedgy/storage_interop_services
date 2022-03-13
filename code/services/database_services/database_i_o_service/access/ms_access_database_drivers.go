package access

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/access/internal"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-adodb"
	"strings"
)

type MsAccessDatabaseDrivers struct {
	OleDb12ConnectionString string
}

func (ms_access_driver *MsAccessDatabaseDrivers) Open() (*sqlx.DB, error) {

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

			database_encoding, database_encoding_error := internal.ReadEncoding(connection_string_variable_value_pair[1])

			if database_encoding_error != nil {
				return nil, database_encoding_error
			}

			connection_string = connection_string + internal.OleDb12ConnectionString_Suffix + database_encoding + `;`
		}
	}

	database, database_open_error :=
		sqlx.Open(internal.DatabaseDriverType, connection_string)

	if database_open_error != nil {
		return nil, database_open_error
	}

	database_open_error = database.Ping()

	if database_open_error != nil {
		return nil, database_open_error
	}

	return database, nil

}
