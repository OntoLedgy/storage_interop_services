package code

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/graph_processors"
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/triples_analysers"
	"github.com/OntoLedgy/storage_interop_services/code/services/rdf_services/triples_processors"
)

func OrchestrateRDFGraphMergeAnalysis(
	inputFile1Name,
	inputFile2Name,
	outputRdfFileName,
	delimiterString string) {

	//Analyse

	inputTripleGraphs,
		mergedDifferenceGraph :=
		triples_analysers.
			AnalyseTriples(
				inputFile1Name,
				inputFile2Name,
				delimiterString)

	//Prepare Mapping Inputs

	mappingGraph :=
		triples_processors.PrepareMappingGraph(
			outputRdfFileName,
			delimiterString)

	//Merge Mapping with difference graph

	graph_processors.
		MergeMappingToDifferenceGraph(
			mergedDifferenceGraph,
			mappingGraph)

	//Merge mapping and input data
	graph_processors.
		MergeAll(
			inputTripleGraphs,
			mappingGraph)
}
