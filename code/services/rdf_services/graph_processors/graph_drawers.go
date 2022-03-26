package graph_processors

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/helpers"
	"github.com/emicklei/dot"
	"github.com/wallix/triplestore"
	"io/ioutil"
)

func GenerateInputTriplesGraphs(
	tripleStores []triplestore.Source) []triplestore.RDFGraph {

	var tripleGraphs []triplestore.RDFGraph

	graph1 :=
		tripleStores[0].
			Snapshot()
	graph2 :=
		tripleStores[1].
			Snapshot()

	DrawGraph(
		graph1,
		"graph_1.dot")
	DrawGraph(
		graph2,
		"graph_2.dot")

	tripleGraphs = append(
		tripleGraphs,
		graph1,
		graph2)

	return tripleGraphs
}

func DrawGraph(
	graph triplestore.RDFGraph,
	filename string) {

	g := dot.NewGraph(dot.Directed)

	triples := graph.Triples()

	for _, triple := range triples {

		subject_node := g.Node(triple.Subject())
		resource_string, _ := triple.Object().Resource()
		object_node := g.Node(resource_string)

		g.Edge(subject_node, object_node, triple.Predicate())

	}

	fmt.Println(g.String())

	err := ioutil.WriteFile("./outputs/"+filename, []byte(g.String()), 0644)
	helpers.Check_error(err)

}
