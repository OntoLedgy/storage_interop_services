package tagger

import (
	"database/sql"
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMastermind_GenerateTag(t *testing.T) {
	type test struct {
		desc     string
		settings func() *configurations.DatabaseToGoSettings
		column   object_model.Column
		expected string
	}

	tests := map[configurations.DatabaseType][]test{
		configurations.DbTypePostgresql: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name: "column_name",
					ConstraintType: sql.NullString{
						String: "PRIMARY KEY",
						Valid:  true,
					},
				},
				expected: `stbl:"column_name,PRIMARY_KEY"`,
			},
			{
				desc: "PK and AI column generates Mastermind-tag with PK and AI indicator",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypePostgresql
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name: "column_name",
					ConstraintType: sql.NullString{
						String: "PRIMARY KEY",
						Valid:  true,
					},
					DefaultValue: sql.NullString{
						String: "nextval",
						Valid:  true,
					},
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
		configurations.DbTypeMySQL: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name:      "column_name",
					ColumnKey: "PRI",
				},
				expected: `stbl:"column_name,PRIMARY_KEY"`,
			},
			{
				desc: "PK and AI column generates Mastermind-tag with PK and AI indicator",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypeMySQL
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name:      "column_name",
					ColumnKey: "PRI",
					Extra:     "auto_increment",
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
		configurations.DbTypeSQLite: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypeSQLite
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name: "column_name",
				},
				expected: `stbl:"column_name"`,
			},
			{
				desc: "PK column generates Mastermind-tag with PK indicator and AI indicator",
				settings: func() *configurations.DatabaseToGoSettings {
					s := configurations.CreateNewSettings()
					s.DbType = configurations.DbTypeSQLite
					s.TagsNoDb = true
					s.TagsMastermindStructable = true
					return s
				},
				column: object_model.Column{
					Name:      "column_name",
					ColumnKey: "PK",
				},
				expected: `stbl:"column_name,PRIMARY_KEY,SERIAL,AUTO_INCREMENT"`,
			},
		},
	}

	tagger := new(Mastermind)

	for dbType := range configurations.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {
			tests := tests[dbType]
			for _, test := range tests {
				t.Run(test.desc, func(t *testing.T) {
					db := database.New(test.settings().Settings)
					actual := tagger.GenerateTag(db, test.column)
					assert.Equal(t, test.expected, actual)
				})
			}
		})
	}
}
