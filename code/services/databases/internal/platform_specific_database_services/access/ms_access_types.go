package access

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/internal"
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
	*internal.DatabaseTransactions
}

type MSAccessQuery contract.Query

type MSAccessQueryx struct {
	*internal.Queryx
}

type MsAccessSqlStatement struct {
	*sqlx.Stmt
}
