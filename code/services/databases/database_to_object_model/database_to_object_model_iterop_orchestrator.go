package database_to_object_model

import (
	"flag"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/pkg/database"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/database_to_object_model/pkg/output"
	"os"
)

func OrchestrateDatabaseToGoInterOp(databaseName string) {

	//TODO - Add config file here.

	settingsFactory := &configurations.SettingsFactory{
		User:           "ladmin",
		DbName:         databaseName,
		Pswd:           "Numark234",
		Host:           "192.168.0.45",
		Port:           "5432",
		DbType:         "pg",
		PackageName:    databaseName,
		OutputFilePath: "D:\\S\\go\\src\\github.com\\OntoLedgy\\domain_ontologies\\code\\data_models\\" + databaseName + "\\",
		Schema:         "public",
	}

	databaseToGoSettings := &configurations.DatabaseToGoSettings{
		Settings: settingsFactory.Create(),
	}

	Orchestrate_Cli(databaseToGoSettings)

}

// main function to run the transformations
func Orchestrate_Cli(databaseToGoSettings *configurations.DatabaseToGoSettings) {

	if databaseToGoSettings.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := databaseToGoSettings.Verify(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	db := database.New(
		databaseToGoSettings.Settings)

	if err := db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writer := output.NewFileWriter(databaseToGoSettings.OutputFilePath)

	if err := RunDatabaseToGoServices(databaseToGoSettings.Settings, db, writer); err != nil {
		fmt.Printf("run error: %v\n", err)
		os.Exit(1)
	}
}
