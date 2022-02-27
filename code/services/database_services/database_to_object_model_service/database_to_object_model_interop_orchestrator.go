package database_to_object_model_service

import (
	"flag"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_to_object_model_service/output/writer"
	"os"
)

func OrchestrateDatabaseToGoInterOp(
	databaseName string,
	outputFilePath string) {

	//TODO - Add config file here.

	settingsFactory := &configurations2.SettingsFactory{
		User:           "ladmin",
		DbName:         databaseName,
		Pswd:           "Numark234",
		Host:           "192.168.0.45",
		Port:           5432,
		DbType:         "pg",
		PackageName:    databaseName,
		OutputFilePath: outputFilePath + databaseName + "\\",
		Schema:         "public",
	}

	databaseToGoSettings := &configurations2.DatabaseToGoSettings{
		Settings: settingsFactory.Create(),
	}

	Orchestrate_Cli(databaseToGoSettings)

}

// main function to run the transformations
func Orchestrate_Cli(databaseToGoSettings *configurations2.DatabaseToGoSettings) {

	if databaseToGoSettings.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := databaseToGoSettings.Verify(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	db := database_i_o_service.New(
		databaseToGoSettings.Settings)

	if err := db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileWriter := writer.NewFileWriter(databaseToGoSettings.OutputFilePath)

	if err := RunDatabaseToGoServices(databaseToGoSettings.Settings, db, fileWriter); err != nil {
		fmt.Printf("run error: %v\n", err)
		os.Exit(1)
	}
}
