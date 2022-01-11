package testing

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model"
	"testing"
)

func TestDatabaseToGo(t *testing.T) {

	outputFolderPath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\domain_ontologies\\code\\data_models\\data_modelling_tools\\"

	database_to_object_model.OrchestrateDatabaseToGoInterOp(
		"sparx_ea",
		outputFolderPath)

}
