package databases

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"testing"
)

func TestAccess(t *testing.T) {

	tests := []struct {
		desc     string
		settings func() *configurations2.Settings
		expected func(*configurations2.Settings) string
	}{
		{
			desc: "test database connection to empty database",
			settings: func() *configurations2.Settings {
				s := configurations2.CreateNewSettings()
				s.DbType = configurations2.DbTypeAccess
				s.DbFileName = "D:\\S\\go\\src\\github.com\\OntoLedgy\\storage_interop_services\\testing\\data\\empty_nf_ea_com.accdb"
				return s.Settings
			},
			expected: func(s *configurations2.Settings) string {
				return fmt.Sprintf("database_i_o_service type = %v, file name =%v ",
					s.DbType, s.DbFileName)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			database_factory := database_i_o_service.DatabaseFactory{
				DatabaseName:   test.settings().DbFileName,
				DatabaseType:   configurations2.DbTypeAccess,
				SystemDatabase: "C:\\Users\\khanm\\AppData\\Roaming\\Microsoft\\Access\\System.mdw"}

			database := database_factory.New()

			connectError := database.Connect()

			defer database.Close()

			if connectError != nil {
				fmt.Println("failed to connect: ", connectError)
			}
			fmt.Println("done")

			tables, tableError :=
				database.GetTables()

			if tableError != nil {
				fmt.Println("failed to get table : ", tableError)
			}

			for index, table := range tables {
				fmt.Printf("table: %v, table name: %s \n",
					index,
					table.Name)
			}

		})
	}
}
