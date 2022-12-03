package neo4j_services

import (
	"context"
	"github.com/OntoLedgy/storage_interop_services/code/services/graph_database_services/object_model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4J struct {
	DriverContext             neo4j.DriverWithContext
	Context                   context.Context
	UniformResourceIdentifier string
}

func (neo4JInstance *Neo4J) Connect(uri string,
	UserName,
	Password string) {

	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	neo4JInstance.UniformResourceIdentifier = uri

	driver, err := neo4j.NewDriverWithContext(
		neo4JInstance.UniformResourceIdentifier,
		neo4j.BasicAuth(
			UserName,
			Password,
			""))

	if err != nil {
		panic(err)
	}

	neo4JInstance.DriverContext = driver

	// Starting with 5.0, you can control the execution of most driver APIs
	// To keep things simple, we create here a never-cancelling context
	// Read https://pkg.go.dev/context to learn more about contexts
	neo4JInstance.Context =
		context.Background()

}

func (neo4JInstance *Neo4J) InsertNode(
	node *object_model.Nodes) (
	*object_model.Nodes,
	error) {
	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.

	session := neo4JInstance.DriverContext.NewSession(
		neo4JInstance.Context,
		neo4j.SessionConfig{})

	defer session.Close(
		neo4JInstance.Context)

	result, err := session.ExecuteWrite(
		neo4JInstance.Context,
		neo4JInstance.executeCreateNodeTransaction(
			node))

	if err != nil {
		return nil, err
	}
	return result.(*object_model.Nodes), nil
}

func (neo4JInstance *Neo4J) executeCreateNodeTransaction(
	node *object_model.Nodes) neo4j.ManagedTransactionWork {

	return func(tx neo4j.ManagedTransaction) (any, error) {

		records, err := tx.Run(
			neo4JInstance.Context,
			"CREATE ("+
				"n:"+node.Type+
				" { "+
				"id: $id, "+
				"name: $name"+
				" }) RETURN n.id, n.name",
			map[string]any{
				"id":   node.Id,
				"name": node.Name,
			})

		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.

		if err != nil {
			return nil, err
		}
		record, err := records.Single(neo4JInstance.Context)
		if err != nil {
			return nil, err
		}

		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`

		return &object_model.Nodes{
			Id:   record.Values[0].(int64),
			Name: record.Values[1].(string),
		}, nil
	}
}

func (neo4JInstance *Neo4J) InsertRelationship(
	node1, node2 *object_model.Nodes) (
	*object_model.Nodes,
	error) {
	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.

	session := neo4JInstance.DriverContext.NewSession(
		neo4JInstance.Context,
		neo4j.SessionConfig{})

	defer session.Close(
		neo4JInstance.Context)

	result, err := session.ExecuteWrite(
		neo4JInstance.Context,
		neo4JInstance.executeCreateRelationshipTransaction(
			node1,
			node2))

	if err != nil {
		return nil, err
	}
	return result.(*object_model.Nodes), nil
}

func (neo4JInstance *Neo4J) executeCreateRelationshipTransaction(
	node1 *object_model.Nodes,
	node2 *object_model.Nodes) neo4j.ManagedTransactionWork {

	return func(tx neo4j.ManagedTransaction) (any, error) {

		records, err := tx.Run(
			neo4JInstance.Context,
			"MATCH "+
				"(a:Types),"+
				"(b:Types)"+
				"WHERE a.name = '"+
				node1.Name+
				"' AND b.name = '"+
				node2.Name+
				"' CREATE (a)-[r:place1]->(b)"+
				"RETURN r",
			nil)

		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.

		if err != nil {
			return nil, err
		}
		record, err := records.Single(neo4JInstance.Context)
		if err != nil {
			return nil, err
		}

		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`

		return &object_model.Nodes{
			Id:   record.Values[0].(int64),
			Name: record.Values[1].(string),
		}, nil
	}
}
