package access

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/constants"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
)

func NewMsAccessDatabase(
	settings *configurations.Settings,
	systemDatabaseFilePath string) *MsAccessDatabase {

	return &MsAccessDatabase{
		GeneralDatabases: &contract.GeneralDatabases{
			Settings: settings,
			Driver:   constants.DbTypeToDriverMap[settings.DbType],
		},
		SystemDatabaseFileName: systemDatabaseFilePath}

}
