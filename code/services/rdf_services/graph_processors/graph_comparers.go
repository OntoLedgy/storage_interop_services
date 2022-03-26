package graph_processors

import "github.com/wallix/triplestore"

func GetDifferencesBetweenGraphs(
	graphs []triplestore.RDFGraph) []triplestore.RDFGraph {

	graph1 := graphs[0]
	graph2 := graphs[1]

	extraGraph1TriplesGraph :=
		compareGraphPair(
			graph1,
			graph2)

	extraGraph2TriplesGraph :=
		compareGraphPair(
			graph2,
			graph1)

	DrawGraph(
		extraGraph1TriplesGraph,
		"extra_in_1.dot")

	DrawGraph(
		extraGraph2TriplesGraph,
		"extra_in_2.dot")

	var mergedDifferenceGraphs []triplestore.RDFGraph

	mergedDifferenceGraphs =
		append(
			mergedDifferenceGraphs,
			extraGraph1TriplesGraph,
			extraGraph2TriplesGraph)

	return mergedDifferenceGraphs

}

func compareGraphPair(
	graph1 triplestore.RDFGraph,
	graph2 triplestore.RDFGraph) triplestore.RDFGraph {

	graph1Triples :=
		graph1.
			Triples()

	graph1TriplesNotInGraph2 :=
		triplestore.
			NewSource()

	for _, graph1Triple := range graph1Triples {

		if !graph2.Contains(graph1Triple) {
			graph1TriplesNotInGraph2.Add(graph1Triple)
		}
	}
	graph1TriplesNotInGraph2Graph := graph1TriplesNotInGraph2.Snapshot()

	return graph1TriplesNotInGraph2Graph
}
