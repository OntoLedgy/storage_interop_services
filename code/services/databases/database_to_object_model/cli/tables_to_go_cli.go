package cli

import (
	"flag"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/object_model/configurations"
)

// DatabaseToGoSettings represents the supported command line args

// NewCmdArgs creates and prepares the command line arguments with default values
func NewCmdArgs() (args *configurations.DatabaseToGoSettings) {

	settingsFactory := &configurations.SettingsFactory{}

	args = &configurations.DatabaseToGoSettings{
		Settings: settingsFactory.Create(),
	}

	flag.BoolVar(&args.Help, "?", false, "shows help and usage")
	flag.BoolVar(&args.Help, "help", false, "shows help and usage")
	flag.BoolVar(&args.Verbose, "v", args.Verbose, "verbose output")
	flag.BoolVar(&args.VVerbose, "vv", args.VVerbose, "more verbose output")
	flag.BoolVar(&args.Force, "f", args.Force, "force; skip tables that encounter errors")

	flag.Var(&args.DbType, "t", fmt.Sprintf("type of database to use, currently supported: %v", configurations.SprintfSupportedDbTypes()))
	flag.StringVar(&args.User, "u", args.User, "user to connect to the database")
	flag.StringVar(&args.Pswd, "p", args.Pswd, "password of user")
	flag.StringVar(&args.DbName, "d", args.DbName, "database name")
	flag.StringVar(&args.Schema, "s", args.Schema, "schema name")
	flag.StringVar(&args.Host, "h", args.Host, "host of database")
	flag.StringVar(&args.Port, "port", args.Port, "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.StringVar(&args.OutputFilePath, "of", args.OutputFilePath, "output file path, default is current working directory")
	flag.Var(&args.OutputFormat, "format", "format of struct fields (columns): camelCase (c) or original (o)")

	flag.Var(&args.FileNameFormat, "fn-format", "format of the filename: camelCase (c, default) or snake_case (s)")
	flag.StringVar(&args.Prefix, "pre", args.Prefix, "prefix for file- and struct names")
	flag.StringVar(&args.Suffix, "suf", args.Suffix, "suffix for file- and struct names")
	flag.StringVar(&args.PackageName, "pn", args.PackageName, "package name")
	flag.Var(&args.Null, "null", "representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)")

	flag.BoolVar(&args.NoInitialism, "no-initialism", args.NoInitialism, "disable the conversion to upper-case words in column names")

	flag.BoolVar(&args.TagsNoDb, "tags-no-db", args.TagsNoDb, "do not create db-tags")

	flag.BoolVar(&args.TagsMastermindStructable, "tags-structable", args.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&args.TagsMastermindStructableOnly, "tags-structable-only", args.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&args.IsMastermindStructableRecorder, "structable-recorder", args.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	// disable the print of usage when an error occurs
	flag.CommandLine.Usage = func() {}

	flag.Parse()

	return args
}
