package database_to_object_model

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/pkg/settings"
)

func OrchestratePostGreToGoInterOp() {

	cmdArgs := &CmdArgs{
		Settings: settings.New(),
	}

	//TODO - Add config file here.

	cmdArgs.User = "ladmin"
	cmdArgs.DbName = "discogs"
	cmdArgs.DbType = ""
	cmdArgs.Pswd = "Numark234"
	cmdArgs.Host = "192.168.0.45"
	cmdArgs.Port = "5432"
	cmdArgs.DbType = "pg"

	cmdArgs.OutputFilePath = "D:\\S\\go\\src\\github.com\\OntoLedgy\\domain_ontologies\\code\\data_models"

	Orchestrate_Cli(cmdArgs)

}
