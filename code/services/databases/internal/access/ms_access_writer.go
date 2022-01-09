package access

/*
import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateMsAccessTable(ms_access_database MsAccessDatabaseDrivers, table_name string, schema AccessSchema) {

	var sql_string string

	var context = context.Background()

	sql_string = create_sql_prefix + table_name + "(" + schema.Schema_string + ");"

	sql_statement, err := ms_access_database.Database.PrepareContext(context, sql_string)

	if err != nil {
		log.Fatal(err)
	}

	_, creation_error := sql_statement.Exec()

	if creation_error != nil {

		fmt.Println(creation_error)

	}

	fmt.Println(creation_error)

	return
}

func Append_to_ms_access_table(ms_access_database *sql.DB, table_name string, dataset [][]interface{}) {

	return
}*/
