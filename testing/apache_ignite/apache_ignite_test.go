package apache_ignite

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/distributed_filesystem_services/apache_ignite_services/connectors"
	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	"log"
	"net"
	"testing"
	"time"
)

var connectionInfo = &ignite.ConnInfo{
	"tcp",
	"192.168.0.30",
	10800,
	1,
	1,
	0,
	"",
	"",
	net.Dialer{
		Timeout: 10 * time.Second},
	nil}

func TestDataApacheIgnite(t *testing.T) {

	connectors.ConnectToIgniteCluster(*connectionInfo)

	println("hello")
}

func TestConnection(t *testing.T) {
	//ctx := context.Background()

	// connect
	c, err := ignite.Connect(*connectionInfo)

	if err != nil {
		t.Fatalf("failed connect to server: %v", err)
	}
	defer c.Close()

	cache := "SQL_PUBLIC_PERSON"

	// create cache
	if err = c.CacheGetOrCreateWithName(cache); err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

}

func TestSQL(t *testing.T) {

	//defer c.CacheDestroy(cache)
	c := connectors.ConnectToIgniteCluster(*connectionInfo)
	// insert data
	cache := "NODES"

	//tm := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)

	// create cache
	if err := c.CacheGetOrCreateWithName(cache); err != nil {
		t.Fatalf("failed to create cache: %v", err)
	}

	table_name := "NODES"
	//create_table(table_name, c, cache, t)

	//insert_into_ignite_table(c, cache, t)

	//select data using QuerySQL
	select_data_from_ignite_table(c, cache, table_name, t)

	//select_data_from_ignite_table(c, cache, t)
}

func select_data_from_ignite_table(c ignite.Client, cache string, table_name string, t *testing.T) {

	//complex object based selection
	r, err := c.QuerySQL(cache, false, ignite.QuerySQLData{
		Table:    "TABLE",
		Query:    "SELECT * FROM " + table_name,
		PageSize: 10000,
	})
	if err != nil {
		t.Fatalf("failed query data: %v", err)
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
			"FROM " + table_name,
		QueryArgs: []interface{}{
			int64(2)},
		Timeout:           10000,
		IncludeFieldNames: true,
	})
	if err != nil {
		t.Fatalf("failed query data: %v", err)
	}

	log.Printf("res=%#v", r2.Rows)

	// complex query using QuerySQLFields
	/*
		r2, err := c.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
			PageSize: 10,
			Query: "SELECT " +
				"o.name AS Name, " +
				"o.foundDateTime AS Found, " +
				"p.firstName AS FirstName, " +
				"p.lastName AS LastName, " +
				"p.salary AS Salary " +
				"FROM Person p INNER JOIN Organization o ON p.orgId = o._key " +
				"WHERE o._key = ? " +
				"ORDER BY p.firstName",
			QueryArgs: []interface{}{
				int64(2)},
			Timeout:           10000,
			IncludeFieldNames: true,
		})
		if err != nil {
			t.Fatalf("failed query data: %v", err)
		}
		log.Printf("res=%#v", r2.Rows)//*/
}

func insert_into_ignite_table(c ignite.Client, cache string, t *testing.T) error {
	_, err := c.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: "INSERT INTO Sink2(ID, name) VALUES" +
			"(?, ?)," +
			"(?, ?)," +
			"(?, ?)",
		QueryArgs: []interface{}{
			int64(1000), "Org 1",
			int64(2000), "Org 2",
			int64(3000), "Org 3"},
	})
	if err != nil {
		t.Fatalf("failed insert data: %v", err)
	}
	return err
}
