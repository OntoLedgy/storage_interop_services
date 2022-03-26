package triples_processors

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/graph_processors"
	"github.com/wallix/triplestore"
)

func GetTriplesGraphs(
	file1Name,
	file2Name,
	delimiterString string) []triplestore.RDFGraph {

	triple_stores :=
		GetTriplesDatasets(
			file1Name,
			file2Name,
			delimiterString)

	triple_input_graphs :=
		graph_processors.GenerateInputTriplesGraphs(
			triple_stores)

	return triple_input_graphs
}

func PrepareMappingGraph(fileName string, delimiter string) triplestore.RDFGraph {

	mapping_dataset :=
		csv.
			ReadCsvToSlice(fileName, delimiter)

	mapping_store :=
		getTripleDataset(
			mapping_dataset)

	mapping_graph :=
		mapping_store.
			Snapshot()

	return mapping_graph
}
