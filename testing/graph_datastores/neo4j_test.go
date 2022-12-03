package graph_datastores

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/graph_database_services/neo4j_services"
	"github.com/OntoLedgy/storage_interop_services/code/services/graph_database_services/object_model"
	"testing"
)

func TestNeo4j(t *testing.T) {
	testNeo4j()
}

func testNeo4j() {

	neo4JInstance := &neo4j_services.Neo4J{}

	neo4JInstance.Connect(
		"neo4j://hitchcock:7687",
		"neo4j",
		"Numark234")

	//dbUri := "neo4j://hitchcock:7687"

	//driver, err, ctx := neo4j_services.Connect()

	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application
	// Make sure to handle errors during deferred calls
	defer neo4JInstance.DriverContext.Close(neo4JInstance.Context)

	testNode1 := &object_model.Nodes{
		Name: "tuple425",
		Id:   12345,
		Type: "Types",
	}
	testNode2 := &object_model.Nodes{
		Name: "type420",
		Id:   12345,
		Type: "Types",
	}

	neo4jnode1, err := neo4JInstance.InsertNode(
		testNode1)

	neo4jnode2, err := neo4JInstance.InsertNode(
		testNode2)

	neo4JInstance.InsertRelationship(neo4jnode1, neo4jnode2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", testNode1)

}
