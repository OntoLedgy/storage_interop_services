package databases

import (
	"fmt"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/postgresql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresql_DSN(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *configurations2.Settings
		expected func(*configurations2.Settings) string
	}{
		{
			desc: "no username given, defaults to `postgres`",
			settings: func() *configurations2.Settings {
				s := configurations2.CreateNewSettings()
				s.DbType = configurations2.DbTypePostgresql
				return s.Settings
			},
			expected: func(s *configurations2.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "postgres", s.DbName, s.Password)
			},
		},
		{
			desc: "with given username, default gets overwritten",
			settings: func() *configurations2.Settings {
				s := configurations2.CreateNewSettings()
				s.DbType = configurations2.DbTypePostgresql
				s.User = "my_custom_user"
				return s.Settings
			},
			expected: func(s *configurations2.Settings) string {
				return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
					s.Host, s.Port, "my_custom_user", s.DbName, s.Password)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := postgresql.NewPostgresql(test.settings())
			actual := db.DSN()
			assert.Equal(t, test.expected(db.Settings), actual)
		})
	}
}
