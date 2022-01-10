package configurations

import (
	"path/filepath"
)

type SettingsFactory struct {
	Verbose  bool
	VVerbose bool
	Force    bool // continue through errors

	DbType DatabaseType

	User   string
	Pswd   string
	DbName string
	Schema string
	Host   string
	Port   string

	OutputFilePath string
	OutputFormat   OutputFormat

	FileNameFormat FileNameFormat
	PackageName    string
	Prefix         string
	Suffix         string
	Null           NullType

	NoInitialism bool

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool

	// TODO not implemented yet
	TagsGorm bool
}

// CreateSettings constructs settings with default values
// TODO - need to add mode - cli vs api or config_file

func (settingsFactory *SettingsFactory) Create() *Settings {

	outputDirectoryPath, err := filepath.Abs(filepath.Dir(settingsFactory.OutputFilePath))

	if err != nil {
		outputDirectoryPath = "."
	}

	packageNameIsNull := settingsFactory.PackageName == ""

	packageName := settingsFactory.PackageName

	if packageNameIsNull {

		packageName = "default_package"
	}

	settings := &Settings{
		Verbose:  false,
		VVerbose: false,
		Force:    false,

		DbType:         settingsFactory.DbType,
		User:           settingsFactory.User,
		Pswd:           settingsFactory.Pswd,
		DbName:         settingsFactory.DbName,
		Schema:         settingsFactory.Schema,
		Host:           settingsFactory.Host,
		Port:           settingsFactory.Port, // left blank -> is automatically determined if not set
		OutputFilePath: outputDirectoryPath,
		OutputFormat:   OutputFormatCamelCase,
		FileNameFormat: FileNameFormatCamelCase,
		PackageName:    packageName,
		Prefix:         "",
		Suffix:         "",
		Null:           NullTypeSQL,

		NoInitialism: false,

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		//TODO to be implemented gorm
		TagsGorm: false,
	}

	return settings
}
