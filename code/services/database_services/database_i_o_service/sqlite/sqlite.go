package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/constants"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	object_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"net/url"
	"strings"
)

// SQLite implemenmts the Database interface with help of generalDatabase
type SQLite struct {
	*contract.GeneralDatabases
}

// NewSQLite creates a new SQLite database_i_o_service
func NewSQLite(s *configurations.Settings) *SQLite {
	return &SQLite{
		GeneralDatabases: &contract.GeneralDatabases{
			Settings: s,
			Driver:   constants.DbTypeToDriverMap[s.DbType],
		},
	}
}

func (s *SQLite) Connect() (err error) {
	return s.GeneralDatabases.Connect(s.DSN())
}

func (s *SQLite) DSN() string {
	if s.Settings.User == "" && s.Settings.Password == "" {
		return fmt.Sprintf("%v", s.Settings.DbName)
	}

	u, err := url.Parse(s.DbName)
	if err != nil {
		return fmt.Sprintf("%v", s.Settings.DbName)
	}

	query := u.Query()
	query.Set("_auth_user", s.Settings.User)
	query.Set("_auth_pass", s.Settings.Password)
	u.RawQuery = query.Encode()

	// SQLite Driver expects a empty `_auth` request param
	return strings.ReplaceAll(u.RequestURI(), "_auth=&", "_auth&")
}

func (s *SQLite) GetDriverImportLibrary() string {
	return `"github.com/mattn/go-sqlite3"`
}

func (s *SQLite) GetTables() (tables []*object_model2.Table, err error) {

	err = s.Select(&tables, `
		SELECT name AS table_name
		FROM sqlite_master
		WHERE type = 'table'
		AND name NOT LIKE 'sqlite?_%' escape '?'
	`)

	if s.Verbose {
		if err != nil {
			fmt.Println("> Error at GetTables()")
			fmt.Printf("> database_i_o_service: %q\r\n", s.DbName)
		}
	}

	return tables, err
}

func (s *SQLite) PrepareGetColumnsOfTableStmt() (err error) {
	return nil
}

func (s *SQLite) GetColumnsOfTable(table *object_model2.Table) (err error) {

	rows, err := s.Queryx(`
		SELECT * 
		FROM PRAGMA_TABLE_INFO('` + table.Name + `')
	`)
	if err != nil {
		if s.Verbose {
			fmt.Printf("> Error at GetColumnsOfTable(%v)\r\n", table.Name)
			fmt.Printf("> database_i_o_service: %q\r\n", s.DbName)
		}
		return err
	}

	type column struct {
		CID          int            `db:"cid"`
		Name         string         `db:"name"`
		DataType     string         `db:"type"`
		NotNull      int            `db:"notnull"`
		DefaultValue sql.NullString `db:"dflt_value"`
		PrimaryKey   int            `db:"pk"`
	}

	for rows.Next() {
		var col column
		err = rows.StructScan(&col)
		if err != nil {
			return err
		}

		isNullable := "YES"
		if col.NotNull == 1 {
			isNullable = "NO"
		}

		isPrimaryKey := ""
		if col.PrimaryKey == 1 {
			isPrimaryKey = "PK"
		}

		table.Columns = append(table.Columns, object_model2.Column{
			OrdinalPosition:        col.CID,
			Name:                   col.Name,
			DataType:               col.DataType,
			DefaultValue:           col.DefaultValue,
			IsNullable:             isNullable,
			CharacterMaximumLength: sql.NullInt64{},
			NumericPrecision:       sql.NullInt64{},
			// reuse mysql column_key as primary key indicator
			ColumnKey:      isPrimaryKey,
			Extra:          "",
			ConstraintName: sql.NullString{},
			ConstraintType: sql.NullString{},
		})
	}

	return nil
}

func (s *SQLite) IsPrimaryKey(column object_model2.Column) bool {
	return column.ColumnKey == "PK"
}

func (s *SQLite) IsAutoIncrement(column object_model2.Column) bool {
	return column.ColumnKey == "PK"
}

func (s *SQLite) GetStringDatatypes() []string {
	return []string{
		"text",
	}
}

func (s *SQLite) IsString(column object_model2.Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetStringDatatypes())
}

func (s *SQLite) GetTextDatatypes() []string {
	return []string{
		"text",
	}
}

func (s *SQLite) IsText(column object_model2.Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetTextDatatypes())
}

func (s *SQLite) GetIntegerDatatypes() []string {
	return []string{
		"integer",
	}
}

func (s *SQLite) IsInteger(column object_model2.Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetIntegerDatatypes())
}

func (s *SQLite) GetFloatDatatypes() []string {
	return []string{
		"real",
		"numeric",
	}
}

func (s *SQLite) IsFloat(column object_model2.Column) bool {
	return s.IsStringInSlice(column.DataType, s.GetFloatDatatypes())
}

func (s *SQLite) GetTemporalDatatypes() []string {
	return []string{}
}

func (s *SQLite) IsTemporal(column object_model2.Column) bool {
	//TODO fix this, column not used.
	return false
}

func (s *SQLite) GetTemporalDriverDataType() string {
	return ""
}
