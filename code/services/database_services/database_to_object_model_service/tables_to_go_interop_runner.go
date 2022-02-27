package database_to_object_model_service

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	object_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_to_object_model_service/output/writer"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_to_object_model_service/tagger"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/language"

	"golang.org/x/text/cases"
	"strings"
	"unicode"
)

var (
	taggers tagger.Tagger

	// some strings for idiomatic go in column names
	// see https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	initialisms = []string{"ID", "JSON", "XML", "HTTP", "URL"}
)

// RunDatabaseToGoServices runs the transformations by creating the concrete Database by the provided settings
func RunDatabaseToGoServices(
	settings *configurations2.Settings,
	sourceDatabase contract.IDatabases,
	out writer.Writer) (
	err error) {

	taggers =
		tagger.NewTaggers(settings)

	fmt.Printf("running for %q...\r\n", settings.DbType)

	sourceDatabaseTables, err :=
		sourceDatabase.GetTables()

	if err != nil {
		return fmt.Errorf("could not get sourceDatabaseTables: %v", err)
	}

	if settings.Verbose {
		fmt.Printf("> number of sourceDatabaseTables: %v\r\n", len(sourceDatabaseTables))
	}

	err = RunTableToGoServices(
		settings,
		sourceDatabase,
		out,
		sourceDatabaseTables)

	if err != nil {
		return err
	}

	fmt.Println("done!")

	return nil
}

func RunTableToGoServices(
	settings *configurations2.Settings,
	sourceDatabase contract.IDatabases,
	out writer.Writer,
	sourceDatabaseTables []*object_model2.Table) error {

	if err := sourceDatabase.PrepareGetColumnsOfTableStmt(); err != nil {
		return fmt.Errorf("could not prepare the get-column-statement: %v", err)
	}

	for _, sourceDatabaseTable := range sourceDatabaseTables {

		if settings.Verbose {
			fmt.Printf("> processing sourceDatabaseTable %q\r\n", sourceDatabaseTable.Name)
		}

		if err := sourceDatabase.GetColumnsOfTable(sourceDatabaseTable); err != nil {
			if !settings.Force {
				return fmt.Errorf("could not get columns of sourceDatabaseTable %q: %v", sourceDatabaseTable.Name, err)
			}
			fmt.Printf("could not get columns of sourceDatabaseTable %q: %v\n", sourceDatabaseTable.Name, err)
			continue
		}

		if settings.Verbose {
			fmt.Printf("\t> number of columns: %v\r\n", len(sourceDatabaseTable.Columns))
		}

		tableName, content, err :=
			createTableStructString(
				settings,
				sourceDatabase,
				sourceDatabaseTable)

		if err != nil {
			if !settings.Force {
				return fmt.Errorf("could not create string for sourceDatabaseTable %q: %v", sourceDatabaseTable.Name, err)
			}
			fmt.Printf("could not create string for sourceDatabaseTable %q: %v\n", sourceDatabaseTable.Name, err)
			continue
		}

		fileName := tableName
		//CamelCaseString(

		if settings.IsFileNameFormatSnakeCase() {
			fileName = strcase.ToSnake(fileName)
		}

		err = out.Write(
			fileName,
			content)

		if err != nil {
			if !settings.Force {
				return fmt.Errorf("could not write struct for sourceDatabaseTable %q: %v", sourceDatabaseTable.Name, err)
			}
			fmt.Printf("could not write struct for sourceDatabaseTable %q: %v\n", sourceDatabaseTable.Name, err)
		}
	}
	return nil
}

func createTableStructString(
	settings *configurations2.Settings,
	sourceDatabase contract.IDatabases,
	sourceDatabaseTable *object_model2.Table) (
	string,
	string,
	error) {

	var structFields strings.Builder
	titleCaser := cases.Title(language.English)

	tableName := titleCaser.String(
		settings.Prefix +
			sourceDatabaseTable.Name +
			settings.Suffix)

	// Replace any whitespace with underscores
	tableName = strings.Map(ReplaceSpace, tableName)
	if settings.IsOutputFormatCamelCase() {
		tableName = CamelCaseString(tableName)
	}

	// Check that the sourceDatabaseTable name doesn't contain any invalid characters for Go variables
	if !ValidVariableName(tableName) {
		return "", "", fmt.Errorf("sourceDatabaseTable name %q contains invalid characters", sourceDatabaseTable.Name)
	}

	columnInfo := object_model2.Columns{}
	columns := map[string]struct{}{}

	for _, column := range sourceDatabaseTable.Columns {

		columnName, err := FormatColumnName(
			settings,
			column.Name,
			sourceDatabaseTable.Name)

		if err != nil {
			return "", "", err
		}

		// ISSUE-4: if columns are part of multiple constraints
		// then the sql returns multiple rows per column name.
		// Therefore we check if we already added a column with
		// that name to the struct, if so, skip.

		if _, ok := columns[columnName]; ok {
			continue
		}
		columns[columnName] = struct{}{}

		if settings.VVerbose {
			fmt.Printf("\t\t> %v\r\n", column.Name)
		}

		columnType, col := mapDbColumnTypeToGoType(settings, sourceDatabase, column)

		// save that we saw types of columns at least once
		if !columnInfo.IsTemporal {
			columnInfo.IsTemporal = col.IsTemporal
		}
		if !columnInfo.IsNullableTemporal {
			columnInfo.IsNullableTemporal = col.IsNullableTemporal
		}
		if !columnInfo.IsNullablePrimitive {
			columnInfo.IsNullablePrimitive = col.IsNullablePrimitive
		}

		structFields.WriteString(columnName)
		structFields.WriteString(" ")
		structFields.WriteString(columnType)
		structFields.WriteString(" ")
		structFields.WriteString(taggers.GenerateTag(sourceDatabase, column))
		structFields.WriteString("\n")
	}

	if settings.IsMastermindStructableRecorder {
		structFields.WriteString("\t\nstructable.Recorder\n")
	}

	var fileContent strings.Builder

	// write header infos
	fileContent.WriteString("package ")
	fileContent.WriteString(settings.PackageName)
	fileContent.WriteString("\n\n")

	// write imports
	generateImports(&fileContent, settings, sourceDatabase, columnInfo)

	// write struct with fields
	fileContent.WriteString("type ")
	fileContent.WriteString(tableName)
	fileContent.WriteString(" struct {\n")
	fileContent.WriteString(structFields.String())
	fileContent.WriteString("}")

	return tableName, fileContent.String(), nil
}

func generateImports(
	content *strings.Builder,
	settings *configurations2.Settings,
	db contract.IDatabases,
	columnInfo object_model2.Columns) {

	if !columnInfo.HasTrue() && !settings.IsMastermindStructableRecorder {
		return
	}

	content.WriteString("import (\n")

	if columnInfo.IsNullablePrimitive && settings.IsNullTypeSQL() {
		content.WriteString("\t\"database_i_o_service/sql\"\n")
	}

	if columnInfo.IsTemporal {
		content.WriteString("\t\"time\"\n")
	}

	if columnInfo.IsNullableTemporal && settings.IsNullTypeSQL() {
		content.WriteString("\t\n")
		content.WriteString(db.GetDriverImportLibrary())
		content.WriteString("\n")
	}

	if settings.IsMastermindStructableRecorder {
		content.WriteString("\t\n\"github.com/Masterminds/structable\"\n")
	}

	content.WriteString(")\n\n")
}

func mapDbColumnTypeToGoType(
	s *configurations2.Settings,
	db contract.IDatabases,
	column object_model2.Column) (
	goType string,
	columnInfo object_model2.Columns) {
	if db.IsString(column) || db.IsText(column) {
		goType = "string"
		if db.IsNullable(column) {
			goType = getNullType(s, "*string", "sql.NullString")
			columnInfo.IsNullable = true
		}
	} else if db.IsInteger(column) {
		goType = "int"
		if db.IsNullable(column) {
			goType = getNullType(s, "*int", "sql.NullInt64")
			columnInfo.IsNullable = true
		}
	} else if db.IsFloat(column) {
		goType = "float64"
		if db.IsNullable(column) {
			goType = getNullType(s, "*float64", "sql.NullFloat64")
			columnInfo.IsNullable = true
		}
	} else if db.IsTemporal(column) {
		if !db.IsNullable(column) {
			goType = "time.Time"
			columnInfo.IsTemporal = true
		} else {
			goType = getNullType(s, "*time.Time", db.GetTemporalDriverDataType())
			columnInfo.IsTemporal = s.Null == configurations2.NullTypeNative
			columnInfo.IsNullableTemporal = true
			columnInfo.IsNullable = true
		}
	} else {
		// TODO handle special data types
		switch column.DataType {
		case "boolean":
			goType = "bool"
			if db.IsNullable(column) {
				goType = getNullType(s, "*bool", "sql.NullBool")
				columnInfo.IsNullable = true
			}
		default:
			goType = getNullType(s, "*string", "sql.NullString")
		}
	}

	columnInfo.IsNullablePrimitive = columnInfo.IsNullable && !db.IsTemporal(column)

	return goType, columnInfo
}

//TODO move this to string_editor_services helper

func CamelCaseString(
	s string) string {

	caseTitle := cases.Title(language.English)

	if s == "" {
		return s
	}

	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		titleCaseString := caseTitle.String(s)

		return titleCaseString

	}

	var cc string
	for _, part := range splitted {
		cc += caseTitle.String(strings.ToLower(part))
	}
	return cc
}

func getNullType(
	settings *configurations2.Settings,
	primitive string,
	sql string) string {
	if settings.IsNullTypeSQL() {
		return sql
	}
	return primitive
}

func ToInitialisms(
	s string) string {
	for _, substr := range initialisms {
		idx := indexCaseInsensitive(s, substr)
		if idx == -1 {
			continue
		}
		toReplace := s[idx : idx+len(substr)]
		s = strings.ReplaceAll(s, toReplace, substr)
	}
	return s
}

func indexCaseInsensitive(
	s, substr string) int {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Index(s, substr)
}

// ValidVariableName checks for the existence of any characters
// outside of Unicode letters, numbers and underscore.
func ValidVariableName(
	s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			return false
		}
	}
	return true
}

// ReplaceSpace swaps any Unicode space characters for underscores
// to create valid Go identifiers
func ReplaceSpace(
	r rune) rune {
	if unicode.IsSpace(r) || r == '\u200B' {
		return '_'
	}
	return r
}

// FormatColumnName checks for invalid characters and transforms a column name
// according to the provided settings.
func FormatColumnName(
	settings *configurations2.Settings,
	column,
	table string) (
	string,
	error) {

	// Replace any whitespace with underscores
	columnName := strings.Map(ReplaceSpace, column)
	columnNameTitleCaser := cases.Title(language.English)

	columnName = columnNameTitleCaser.String(columnName)

	if settings.IsOutputFormatCamelCase() {
		columnName = CamelCaseString(columnName)
	}
	if settings.ShouldInitialism() {
		columnName = ToInitialisms(columnName)
	}

	// Check that the column name doesn't contain any invalid characters for Go variables
	if !ValidVariableName(columnName) {
		return "", fmt.Errorf("column name %q in table %q contains invalid characters", column, table)
	}
	// First character of an identifier in Go must be letter or _
	// We want it to be an uppercase letter to be a public field
	if !unicode.IsLetter([]rune(columnName)[0]) {
		prefix := "X_"
		if settings.IsOutputFormatCamelCase() {
			prefix = "X"
		}
		if settings.Verbose {
			fmt.Printf("\t\t>column %q in table %q doesn't start with a letter; prepending with %q\n", column, table, prefix)
		}
		columnName = prefix + columnName
	}

	return columnName, nil
}
