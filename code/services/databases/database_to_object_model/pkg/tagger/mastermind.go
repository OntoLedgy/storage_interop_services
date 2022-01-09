package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/pkg/database"
)

// Mastermind represents the Masterminds/structable "stbl"-tag
type Mastermind struct{}

// GenerateTag for Mastermind to satisfy the Tagger interface
func (t Mastermind) GenerateTag(db database.Database, column database.Column) string {

	isPk := ""
	if db.IsPrimaryKey(column) {
		isPk = ",PRIMARY_KEY"
	}

	isAutoIncrement := ""
	if db.IsAutoIncrement(column) {
		isAutoIncrement = ",SERIAL,AUTO_INCREMENT"
	}

	return `stbl:"` + column.Name + isPk + isAutoIncrement + `"`
}
