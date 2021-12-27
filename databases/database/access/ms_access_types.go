package access

import (
	"database_manager/general_database"
	"github.com/jmoiron/sqlx"
)

type AccessSchema struct {
	Schema_string string

	/*Fields struct {
		field_names []string
		field_data_types []string
	}*/

}

type MSAccessStatement struct {
	*general_database.Tx
}

type MSAccessQuery general_database.Query

type MSAccessQueryx struct {
	*general_database.Queryx
}

type MsAccessSqlStatement struct {
	*sqlx.Stmt
}
