package access

//TODO - Deprecate this.

/*
import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-adodb"
	"strings"
)

//
// Implementing decoding of the so called password
//



func init() {
	//database := MsAccessDatabaseDrivers{}
	//general_database.Register(database_driver_type, &database)
}

//
// Implementing new open function
//



func OpenMsAccessDatabase(
	database_filename string) (
		*MsAccessDatabaseDrivers,
		error) {

	var database_open_error error
	var ms_access_database MsAccessDatabaseDrivers

	connection_string := OleDb12ConnectionString_Prefix + database_filename

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

	ms_access_database.Database, database_open_error = sqlx.Open(database_driver_type, connection_string)

	if database_open_error != nil {
		return nil, database_open_error
	}

	database_open_error = ms_access_database.Database.Ping()

	if database_open_error != nil {
		return nil, database_open_error
	}
	return &ms_access_database, nil
} */
