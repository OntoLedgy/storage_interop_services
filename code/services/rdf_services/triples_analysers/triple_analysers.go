package triples_analysers

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/graph_processors"
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/triples_processors"
	"github.com/wallix/triplestore"
)

func AnalyseTriples(
	file1Name,
	file2Name,
	delimiterString string) (
	[]triplestore.RDFGraph,
	triplestore.RDFGraph) {

	triple_input_graphs :=
		triples_processors.
			GetTriplesGraphs(
				file1Name,
				file2Name,
				delimiterString)

	graph_processors.
		AnalyseGraphsPopulation(
			triple_input_graphs)

	merged_difference_graph :=
		report_differences_between_graphs(
			triple_input_graphs)

	return triple_input_graphs, merged_difference_graph
}

func report_differences_between_graphs(
	triple_input_graphs []triplestore.RDFGraph) triplestore.RDFGraph {

	merged_difference_graph :=
		compare_triples_graphs(
			triple_input_graphs)

	graph_processors.DrawGraph(
		merged_difference_graph,
		"merged_difference_graph.dot")

	return merged_difference_graph
}

func compare_triples_graphs(
	triple_input_graphs []triplestore.RDFGraph) triplestore.RDFGraph {

	difference_graphs :=
		graph_processors.
			GetDifferencesBetweenGraphs(
				triple_input_graphs)

	merged_difference_graph :=
		graph_processors.
			MergeGraphs(
				difference_graphs)

	return merged_difference_graph
}
