package testing

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model"
	"testing"
)

func TestDatabaseToGo(t *testing.T) {

	database_to_object_model.OrchestratePostGreToGoInterOp()

}
