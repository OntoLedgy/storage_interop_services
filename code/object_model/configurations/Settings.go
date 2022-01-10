package configurations

var (
	// SupportedDbTypes represents the supported databases
	SupportedDbTypes = map[DatabaseType]bool{
		DbTypePostgresql: true,
		DbTypeMySQL:      true,
		DbTypeSQLite:     true,
	}

	// supportedOutputFormats represents the supported output formats
	supportedOutputFormats = map[OutputFormat]bool{
		OutputFormatCamelCase: true,
		OutputFormatOriginal:  true,
	}

	// dbDefaultPorts maps the database type to the default ports
	dbDefaultPorts = map[DatabaseType]string{
		DbTypePostgresql: "5432",
		DbTypeMySQL:      "3306",
		DbTypeSQLite:     "",
	}

	// SupportedNullTypes represents the supported types of NULL types
	SupportedNullTypes = map[NullType]bool{
		NullTypeSQL:       true,
		NullTypeNative:    true,
		NullTypePrimitive: true,
	}

	// supportedFileNameFormats represents the supported filename formats
	supportedFileNameFormats = map[FileNameFormat]bool{
		FileNameFormatCamelCase: true,
		FileNameFormatSnakeCase: true,
	}
)

// Settings stores the supported settings / command line arguments
type Settings struct {
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
