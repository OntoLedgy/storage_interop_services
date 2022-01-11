package tagger

import (
	"github.com/OntoLedgy/storage_interop_services/code/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaggers_GenerateTags(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *configurations.DatabaseToGoSettings
		column   object_model.Column
		expected string
	}{
		{
			desc: "enabled db-tag (default) without other tags generates only db-tags",
			settings: func() *configurations.DatabaseToGoSettings {
				s := configurations.CreateNewSettings()
				s.TagsNoDb = false
				return s
			},
			column: object_model.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\"`",
		},
		{
			desc: "disabled db-tag without other tags generates no tags",
			settings: func() *configurations.DatabaseToGoSettings {
				s := configurations.CreateNewSettings()
				s.TagsNoDb = true
				return s
			},
			column:   object_model.Column{},
			expected: "",
		},
		{
			desc: "default db-tag with enabled Mastermind-tag creates db- and Mastermind-tags",
			settings: func() *configurations.DatabaseToGoSettings {
				s := configurations.CreateNewSettings()
				s.TagsNoDb = false
				s.TagsMastermindStructable = true
				return s
			},
			column: object_model.Column{
				Name: "column_name",
			},
			expected: "`db:\"column_name\" stbl:\"column_name\"`",
		},
		{
			desc: "disabled db-tag with enabled Mastermind-tag creates only Mastermind-tags",
			settings: func() *configurations.DatabaseToGoSettings {
				s := configurations.CreateNewSettings()
				s.TagsNoDb = true
				s.TagsMastermindStructable = true
				return s
			},
			column: object_model.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
		{
			desc: "default db-tag with enabled standalone Mastermind-tag creates only standalone Mastermind-tag",
			settings: func() *configurations.DatabaseToGoSettings {
				s := configurations.CreateNewSettings()
				s.TagsNoDb = false
				s.TagsMastermindStructableOnly = true
				return s
			},
			column: object_model.Column{
				Name: "column_name",
			},
			expected: "`stbl:\"column_name\"`",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings()
			taggers := NewTaggers(s.Settings)
			db := database.New(s.Settings)
			actual := taggers.GenerateTag(db, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
