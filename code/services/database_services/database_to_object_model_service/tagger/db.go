package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
)

// Db is the standard "db"-tag
type Db struct{}

// GenerateTag for Db to satisfy the Tagger interface
func (t Db) GenerateTag(db contract.IDatabases, column object_model.Column) string {
	//TODO fix this, db not being used.
	return `db:"` + column.Name + `"`
}
