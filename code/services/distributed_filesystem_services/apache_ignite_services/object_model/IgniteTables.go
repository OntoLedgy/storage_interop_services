package object_model

import (
	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	"testing"
)

func CreateTable(
	table_name string,
	igniteClient ignite.Client,
	cache string,
	t *testing.T) {

	_, err := igniteClient.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: `CREATE TABLE ` +
			table_name +
			` (ID INTEGER PRIMARY KEY, NAME VARCHAR(100));`,
		QueryArgs: []interface{}{}})
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

}
