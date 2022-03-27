package apache_ignite_services

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/distributed_filesystem_services/apache_ignite_services/connectors"
	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	"log"
)

func PrepareIginite(
	info ignite.ConnInfo) {

	c := connectors.ConnectToIgniteCluster(info)

	defer c.Close()

	cache := "SQL_PUBLIC_PERSON"

	// create cache

	if err := c.CacheGetOrCreateWithName(cache); err != nil {
		fmt.Printf("failed to create cache: %v", err)
	}

}

func createIgniteTable(
	c ignite.Client,
	cache string,
	table_name string) {

	_, err := c.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: `CREATE TABLE ` +
			table_name +
			` (ID INTEGER PRIMARY KEY, NAME VARCHAR(100));`,
		QueryArgs: []interface{}{}})

	if err != nil {
		fmt.Printf("failed to create table: %v", err)
	}

}

func insertIntoIgniteTable(
	c ignite.Client,
	cache string,
	table_name string) {

	_, err := c.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: "INSERT INTO " +
			table_name +
			"(ID, name) VALUES" +
			"(?, ?)," +
			"(?, ?)," +
			"(?, ?)",
		QueryArgs: []interface{}{
			int64(1000), "Org 1",
			int64(2000), "Org 2",
			int64(3000), "Org 3"},
	})
	if err != nil {
		fmt.Printf("failed insert data: %v", err)
	}

}

func selectDataFromIgniteTable(
	c ignite.Client,
	cache string) {

	//complex object based selection
	r, err := c.QuerySQL(cache, false, ignite.QuerySQLData{
		Table:    "PERSON",
		Query:    "SELECT * FROM PERSON",
		PageSize: 10000,
	})
	if err != nil {
		fmt.Printf("failed query data: %v", err)
	}
	row := r.Rows[int64(1)].(ignite.ComplexObject)
	log.Printf("%d=\"%s\", %d=%#v", 1, row.Fields[1], 2, row.Fields[2])
	row = r.Rows[int64(2)].(ignite.ComplexObject)
	log.Printf("%d=\"%s\", %d=%#v", 1, row.Fields[1], 2, row.Fields[2])
	row = r.Rows[int64(3)].(ignite.ComplexObject)
	log.Printf("%d=\"%s\", %d=%#v", 1, row.Fields[1], 2, row.Fields[2])

	//simple object selection
	r2, err := c.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: "SELECT *" +
			"FROM + Person",
		QueryArgs: []interface{}{
			int64(2)},
		Timeout:           10000,
		IncludeFieldNames: true,
	})
	if err != nil {
		fmt.Printf("failed query data: %v", err)
	}

	log.Printf("res=%#v", r2.Rows)
}
