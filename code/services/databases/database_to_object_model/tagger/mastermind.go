package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/contract"
)

// Mastermind represents the Masterminds/structable "stbl"-tag
type Mastermind struct{}

// GenerateTag for Mastermind to satisfy the Tagger interface
func (t Mastermind) GenerateTag(db contract.Database, column object_model.Column) string {

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
