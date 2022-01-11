package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
)

// Db is the standard "db"-tag
type Db struct{}

// GenerateTag for Db to satisfy the Tagger interface
func (t Db) GenerateTag(db contract.Database, column object_model.Column) string {
	//TODO fix this, db not being used.
	return `db:"` + column.Name + `"`
}
