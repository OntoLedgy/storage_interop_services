package databases

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresql_DSN(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *configurations.Settings
		expected func(*configurations.Settings) string
	}{
		{
			desc: "no username given, defaults to `postgres`",
			settings: func() *configurations.Settings {
				s := configurations.CreateNewSettings()
				s.DbType = configurations.DbTypePostgresql
				return s.Settings
			},
			expected: func(s *configurations.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "postgres", s.DbName, s.Pswd)
			},
		},
		{
			desc: "with given username, default gets overwritten",
			settings: func() *configurations.Settings {
				s := configurations.CreateNewSettings()
				s.DbType = configurations.DbTypePostgresql
				s.User = "my_custom_user"
				return s.Settings
			},
			expected: func(s *configurations.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "my_custom_user", s.DbName, s.Pswd)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := database.NewPostgresql(test.settings())
			actual := db.DSN()
			assert.Equal(t, test.expected(db.Settings), actual)
		})
	}
}
