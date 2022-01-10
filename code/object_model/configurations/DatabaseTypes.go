package configurations

import "fmt"

// DatabaseType represents a type of a database.
type DatabaseType string

const (
	DbTypePostgresql DatabaseType = "pg"
	DbTypeMySQL      DatabaseType = "mysql"
	DbTypeSQLite     DatabaseType = "sqlite3"
)

// Set sets the datatype for the custom type for the flag package.
func (db *DatabaseType) Set(s string) error {
	*db = DatabaseType(s)
	if *db == "" {
		*db = DbTypePostgresql
	}
	if !SupportedDbTypes[*db] {
		return fmt.Errorf("type of database %q not supported! supported: %v", *db, SprintfSupportedDbTypes())
	}
	return nil
}

// String is the implementation of the Stringer interface needed for flag.Value interface.
func (db DatabaseType) String() string {
	return string(db)
}
