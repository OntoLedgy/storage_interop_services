package testing

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccess(t *testing.T) {

	tests := []struct {
		desc     string
		settings func() *configurations.Settings
		expected func() string
	}{
		{
			desc: "test storage interop services facade - no datastore type",
			settings: func() *configurations.Settings {
				s := configurations.CreateNewSettings()
				s.DbType = configurations.DbTypeAccess
				s.DbFileName = "D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\empty_nf_ea_com.accdb"
				return s.Settings
			},
			expected: func() string {
				return fmt.Sprintf("select storage technology")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			storage_service_facade := &services.StorageServices{}

			actual := storage_service_facade.StoreData("")
			assert.Equal(t, test.expected(), actual.Error())
		})
	}
}
