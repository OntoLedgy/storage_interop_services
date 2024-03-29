package databases

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service"
	object_model2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	configurations2 "github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model/configurations"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_to_object_model_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCaseString(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "empty string returns empty string",
			input:    "",
			expected: "",
		},
		{
			desc:     "single string returns titleized single string",
			input:    "string",
			expected: "String",
		},
		{
			desc:     "multi separated string returns CamelCase string",
			input:    "string_with_separate_sections",
			expected: "StringWithSeparateSections",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := database_to_object_model_service.CamelCaseString(tt.input)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.input)
		})
	}
}

func TestToInitialisms(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "id should be upper case",
			input:    "Id",
			expected: "ID",
		},
		{
			desc:     "id at the end of string should be upper case",
			input:    "userId",
			expected: "userID",
		},
		{
			desc:     "id at the beginning of string should be upper case",
			input:    "Iduser",
			expected: "IDuser",
		},
		{
			desc:     "id in the middle of string should be upper case",
			input:    "userIdprim",
			expected: "userIDprim",
		},
		{
			desc:     "multiple occurrences should be upper case",
			input:    "userIdasJsonWithUrl",
			expected: "userIDasJSONWithURL",
		},
		{
			desc:     "multiple id in the string should be upper case",
			input:    "IduserId",
			expected: "IDuserID",
		},
		{
			desc:     "non replacement in the string should be return original string",
			input:    "name",
			expected: "name",
		},
		{
			desc:     "replacements only in the string should be return original string",
			input:    "IdjsonuRlHtTp",
			expected: "IDJSONURLHTTP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := database_to_object_model_service.ToInitialisms(tt.input)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.input)
		})
	}
}

func TestRun_StringTextColumns(t *testing.T) {
	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := configurations2.CreateNewSettings()
			s.DbType = dbType
			db := database_i_o_service.New(s.Settings)

			columnTypes := db.GetStringDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {

						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package default_package\n\ntype TestTable struct {\nColumnName string `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullString `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *string `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *string `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table1 := &object_model2.Table{
							Name: "test_table_1",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &object_model2.Table{
							Name: "test_table_2",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table1, table2)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							On("GetColumnsOfTable", table2)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullString `db:\"column_name_1\"`\nColumnName2 string `db:\"column_name_2\"`\n}",
							).
							On(
								"Write",
								"TestTable2",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 string `db:\"column_name_1\"`\nColumnName2 sql.NullString `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_IntegerColumns(t *testing.T) {
	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := configurations2.CreateNewSettings()
			s.DbType = dbType
			db := database_i_o_service.New(s.Settings)

			columnTypes := db.GetIntegerDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\ntype TestTable struct {\nColumnName int `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullInt64 `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *int `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullInt64 `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *int `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table1 := &object_model2.Table{
							Name: "test_table_1",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &object_model2.Table{
							Name: "test_table_2",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table1, table2)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							On("GetColumnsOfTable", table2)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullInt64 `db:\"column_name_1\"`\nColumnName2 int `db:\"column_name_2\"`\n}",
							).
							On(
								"Write",
								"TestTable2",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 int `db:\"column_name_1\"`\nColumnName2 sql.NullInt64 `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_FloatColumns(t *testing.T) {
	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := configurations2.CreateNewSettings()
			s.DbType = dbType
			db := database_i_o_service.New(s.Settings)

			columnTypes := db.GetFloatDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\ntype TestTable struct {\nColumnName float64 `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullFloat64 `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *float64 `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullFloat64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *float64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table1 := &object_model2.Table{
							Name: "test_table_1",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &object_model2.Table{
							Name: "test_table_2",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table1, table2)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							On("GetColumnsOfTable", table2)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullFloat64 `db:\"column_name_1\"`\nColumnName2 float64 `db:\"column_name_2\"`\n}",
							).
							On(
								"Write",
								"TestTable2",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 float64 `db:\"column_name_1\"`\nColumnName2 sql.NullFloat64 `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_TemporalColumns(t *testing.T) {
	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := configurations2.CreateNewSettings()
			s.DbType = dbType
			db := database_i_o_service.New(s.Settings)

			columnTypes := db.GetTemporalDatatypes()

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName time.Time `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\n"+db.GetDriverImportLibrary()+"\n)\n\ntype TestTable struct {\nColumnName "+dbType.String()+".NullTime `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName *time.Time `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n\t\n"+db.GetDriverImportLibrary()+"\n)\n\ntype TestTable struct {\nColumnName1 "+dbType.String()+".NullTime `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n)\n\ntype TestTable struct {\nColumnName1 *time.Time `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table1 := &object_model2.Table{
							Name: "test_table_1",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &object_model2.Table{
							Name: "test_table_2",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table1, table2)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							On("GetColumnsOfTable", table2)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n\t\n"+db.GetDriverImportLibrary()+"\n)\n\ntype TestTable1 struct {\nColumnName1 "+dbType.String()+".NullTime `db:\"column_name_1\"`\nColumnName2 time.Time `db:\"column_name_2\"`\n}",
							).
							On(
								"Write",
								"TestTable2",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"time\"\n\t\n"+db.GetDriverImportLibrary()+"\n)\n\ntype TestTable2 struct {\nColumnName1 time.Time `db:\"column_name_1\"`\nColumnName2 "+dbType.String()+".NullTime `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestRun_BooleanColumns(t *testing.T) {
	for dbType := range configurations2.SupportedDbTypes {
		t.Run(dbType.String(), func(t *testing.T) {

			s := configurations2.CreateNewSettings()
			s.DbType = dbType
			db := database_i_o_service.New(s.Settings)

			columnTypes := []string{"boolean"}

			for _, columnType := range columnTypes {
				t.Run(columnType, func(t *testing.T) {

					t.Run("single table with NOT NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\ntype TestTable struct {\nColumnName bool `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName sql.NullBool `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with NULL column and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName *bool `db:\"column_name\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable struct {\nColumnName1 sql.NullBool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("single table with two mixed columns and native data type", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType
						s.Null = configurations2.NullTypeNative

						mdb := newMockDb(db)

						table := &object_model2.Table{
							Name: "test_table",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						mdb.tables = append(mdb.tables, table)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n)\n\ntype TestTable struct {\nColumnName1 *bool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})

					t.Run("multi table with multi columns", func(t *testing.T) {
						s := configurations2.CreateNewSettings()
						s.DbType = dbType

						mdb := newMockDb(db)

						table1 := &object_model2.Table{
							Name: "test_table_1",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
									IsNullable:      "YES",
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
								},
							},
						}
						table2 := &object_model2.Table{
							Name: "test_table_2",
							Columns: []object_model2.Column{
								{
									OrdinalPosition: 1,
									Name:            "column_name_1",
									DataType:        columnType,
								},
								{
									OrdinalPosition: 2,
									Name:            "column_name_2",
									DataType:        columnType,
									IsNullable:      "YES",
								},
							},
						}
						mdb.tables = append(mdb.tables, table1, table2)

						mdb.
							On("GetTables").
							Return(mdb.tables, nil)
						mdb.
							On("PrepareGetColumnsOfTableStmt").
							Return(nil)
						mdb.
							On("GetColumnsOfTable", table1).
							On("GetColumnsOfTable", table2)

						w := newMockWriter()
						w.
							On(
								"Write",
								"TestTable1",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable1 struct {\nColumnName1 sql.NullBool `db:\"column_name_1\"`\nColumnName2 bool `db:\"column_name_2\"`\n}",
							).
							On(
								"Write",
								"TestTable2",
								"package "+object_model2.DefaultPacakgeName+"\n\nimport (\n\t\"database_i_o_service/sql\"\n)\n\ntype TestTable2 struct {\nColumnName1 bool `db:\"column_name_1\"`\nColumnName2 sql.NullBool `db:\"column_name_2\"`\n}",
							)

						err := database_to_object_model_service.RunDatabaseToGoServices(s.Settings, mdb, w)
						assert.NoError(t, err)
					})
				})
			}
		})
	}
}

func TestValidVariableName(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected bool
	}
	tests := []testCase{
		{"basic", "MyVariable_2", true},
		{"specialChars", "MyVar;iable", false},
		{"brackets", "MyVariabl(e)", false},
		{"nonEnglish", "MyVαriαble", true},
		{"spaces", "My Variable", false},
		{"whitespace", "My		Variable", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if database_to_object_model_service.ValidVariableName(tc.input) != tc.expected {
				t.Errorf("TestValidVariableName(%q) should be %t", tc.input, tc.expected)
			}
		})
	}
}

func TestReplaceSpace(t *testing.T) {
	type testCase struct {
		name     string
		input    rune
		expected rune
	}
	tests := []testCase{
		{"letter", 'a', 'a'},
		{"number", '7', '7'},
		{"nonEnglish", '水', '水'},
		{"space", ' ', '_'},
		{"underscore", '_', '_'},
		{"tab", '\t', '_'},
		{"newline", '\n', '_'},
		{"zeroWidthSpace", '​', '_'},
		{"nonBreakingSpace", ' ', '_'},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := database_to_object_model_service.ReplaceSpace(tc.input)
			if output != tc.expected {
				t.Errorf("replaceSpace(%q) = %q, expected %q", tc.input, output, tc.expected)
			}
		})
	}

}

func TestFormatColumnName(t *testing.T) {
	// success and failure subtests
	t.Run("pass", func(t *testing.T) {
		type testCase struct {
			name     string
			input    string
			original string
			camel    string
		}
		tests := []testCase{
			{"startWithNumber", "1fish2fish", "X_1Fish2fish", "X1Fish2fish"},
			{"containsSpaces", "my column\twith\nmany\u200bspaces", "My_column_with_many_spaces", "MyColumnWithManySpaces"},
			{"titleCase", "MyColumn", "Mycolumn", "Mycolumn"},
			{"snakeCase", "my_column", "My_column", "MyColumn"},
			{"titleSnake", "My_Column", "My_column", "MyColumn"},
			{"numbersOnly", "123", "X_123", "X123"},
			{"nonEnglish", "火", "火", "火"},
			{"nonEnglishUpper", "Λλ", "Λλ", "Λλ"},
		}
		// subtests for camelCase and original settings
		camelSettings := configurations2.CreateNewSettings()
		camelSettings.OutputFormat = configurations2.OutputFormatCamelCase
		originalSettings := configurations2.CreateNewSettings()
		originalSettings.OutputFormat = configurations2.OutputFormatOriginal
		t.Run("camelcase", func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					output, err := database_to_object_model_service.FormatColumnName(camelSettings.Settings, tc.input, "MyTable")
					if err != nil {
						t.Error(err)
					} else if output != tc.camel {
						t.Errorf("camelcase format of %q = %q, expected %q", tc.input, output, tc.camel)
					}
				})
			}
		})
		t.Run("original", func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					output, err := database_to_object_model_service.FormatColumnName(originalSettings.Settings, tc.input, "MyTable")
					if err != nil {
						t.Error(err)
					} else if output != tc.original {
						t.Errorf("originalCase format of %q = %q, expected %q", tc.input, output, tc.original)
					}
				})
			}
		})
	})
	t.Run("fail", func(t *testing.T) {
		type testCase struct {
			name  string
			input string
		}
		tests := []testCase{
			{"semicolons", "MyColumn;"},
			{"brackets", "MyColumn()"},
		}
		settings := configurations2.CreateNewSettings()
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				_, err := database_to_object_model_service.FormatColumnName(settings.Settings, tc.input, "MyTable")
				if err == nil {
					t.Errorf("formatColumnName(%q) should have thrown error but didn't", tc.input)
				}
			})
		}
	})
}
