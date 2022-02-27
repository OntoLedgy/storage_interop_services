package configurations

import (
	"fmt"
	"os"
	"path/filepath"
)

type DatabaseToGoSettings struct {
	Help bool
	*Settings
}

// Verify verifies the settings and checks the given output paths
func (settings *Settings) Verify() (err error) {

	if err = settings.verifyOutputPath(); err != nil {
		return err
	}

	if settings.OutputFilePath, err = settings.prepareOutputPath(); err != nil {
		return err
	}

	if settings.Port == 0 {
		settings.Port = dbDefaultPorts[settings.DbType]
	}

	if settings.PackageName == "" {
		return fmt.Errorf("name of package can not be empty")
	}

	if settings.VVerbose {
		settings.Verbose = true
	}

	return err
}

func (settings *Settings) verifyOutputPath() (err error) {

	info, err := os.Stat(settings.OutputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("output file path %q does not exists", settings.OutputFilePath)
	}

	if !info.Mode().IsDir() {
		return fmt.Errorf("output file path %q is not a directory", settings.OutputFilePath)
	}

	return err
}

func (settings *Settings) prepareOutputPath() (outputFilePath string, err error) {
	outputFilePath, err = filepath.Abs(settings.OutputFilePath)
	outputFilePath += string(filepath.Separator)
	return outputFilePath, err
}

// IsNullTypeSQL returns true if the type given by the command line args is of null type SQL
func (settings *Settings) IsNullTypeSQL() bool {
	return settings.Null == NullTypeSQL
}

// ShouldInitialism returns wheather or not if column names should be converted
// to initialisms.
func (settings *Settings) ShouldInitialism() bool {
	return !settings.NoInitialism
}

// IsOutputFormatCamelCase returns if the type given by command line args is of camel-case format.
func (settings *Settings) IsOutputFormatCamelCase() bool {
	return settings.OutputFormat == OutputFormatCamelCase
}

// IsFileNameFormatSnakeCase returns if the type given by the command line args is snake-case format
func (settings *Settings) IsFileNameFormatSnakeCase() bool {
	return settings.FileNameFormat == FileNameFormatSnakeCase
}

// SprintfSupportedDbTypes returns a slice of strings as names of the supported database_i_o_service types
func SprintfSupportedDbTypes() string {
	names := make([]string, 0, len(SupportedDbTypes))
	for name := range SupportedDbTypes {
		names = append(names, string(name))
	}
	return fmt.Sprintf("%v", names)
}

// SprintfSupportedNullTypes returns a slice of strings as names of the supported null types
func SprintfSupportedNullTypes() string {
	names := make([]string, 0, len(SupportedNullTypes))
	for name := range SupportedNullTypes {
		names = append(names, string(name))
	}
	return fmt.Sprintf("%v", names)
}

func CreateNewSettings() *DatabaseToGoSettings {
	settingsFactory := &SettingsFactory{}

	databaseToGoSettings := &DatabaseToGoSettings{
		Settings: settingsFactory.Create(),
	}

	return databaseToGoSettings
}
