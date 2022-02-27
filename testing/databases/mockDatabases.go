package databases

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/contract"
	"github.com/OntoLedgy/storage_interop_services/code/services/database_services/database_i_o_service/object_model"
	"github.com/stretchr/testify/mock"
)

type mockDb struct {
	mock.Mock
	contract.IDatabases

	tables []*object_model.Table
}

func newMockDb(db contract.IDatabases) *mockDb {
	return &mockDb{IDatabases: db}
}

func (db *mockDb) Connect() (err error) {
	db.Called()
	return nil
}

func (db *mockDb) Close() (err error) {
	db.Called()
	return nil
}

func (db *mockDb) GetTables() (tables []*object_model.Table, err error) {
	db.Called()
	return db.tables, nil
}

func (db *mockDb) PrepareGetColumnsOfTableStmt() (err error) {
	db.Called()
	return nil
}

func (db *mockDb) GetColumnsOfTable(table *object_model.Table) (err error) {
	db.Called(table)
	return nil
}

type mockWriter struct {
	mock.Mock
}

func newMockWriter() *mockWriter {
	return &mockWriter{}
}

func (w *mockWriter) Write(tableName string, content string) error {
	w.Called(tableName, content)
	return nil
}
