package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/pkg/database"
)

// Db is the standard "db"-tag
type Db struct{}

// GenerateTag for Db to satisfy the Tagger interface
func (t Db) GenerateTag(db database.Database, column database.Column) string {
	return `db:"` + column.Name + `"`
}
