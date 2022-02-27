package tagger

import (
	"database/sql"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMastermind_GenerateTag(t *testing.T) {
	type test struct {
		desc     string
		settings func() *configurations2.DatabaseToGoSettings
		column   object_model.Column
		expected string
	}

	tests := map[configurations2.DatabaseType][]test{
		configurations2.DbTypePostgresql: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypePostgresql
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
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypePostgresql
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
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypePostgresql
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
		configurations2.DbTypeMySQL: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypeMySQL
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
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypeMySQL
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
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypeMySQL
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
		configurations2.DbTypeSQLite: {
			{
				desc: "non PK column generates standard Mastermind-tag",
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypeSQLite
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
				settings: func() *configurations2.DatabaseToGoSettings {
					s := configurations2.CreateNewSettings()
					s.DbType = configurations2.DbTypeSQLite
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

	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {
			tests := tests[dbType]
			for _, test := range tests {
				t.Run(test.desc, func(t *testing.T) {
					db := contract.New(test.settings().Settings)
					actual := tagger.GenerateTag(db, test.column)
					assert.Equal(t, test.expected, actual)
				})
			}
		})
	}
}
