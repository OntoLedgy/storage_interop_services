package graph_processors

import (
	"github.com/wallix/triplestore"
)

func MergeAll(
	tripleInputGraphs []triplestore.RDFGraph,
	mappingGraph triplestore.RDFGraph) {
	mergedGraphExcludingMapping :=
		MergeGraphs(
			tripleInputGraphs)

	var tripleInputAndMappingGraphs []triplestore.RDFGraph

	tripleInputAndMappingGraphs =
		append(
			tripleInputAndMappingGraphs,
			mergedGraphExcludingMapping,
			mappingGraph)

	mergedGraph :=
		MergeGraphs(
			tripleInputAndMappingGraphs)
	DrawGraph(
		mergedGraph,
		"merged_graph_including_mapping.dot")

	AnalyseGraphPopulation(
		mergedGraph,
		1)

}

func MergeMappingToDifferenceGraph(
	mergedDifferenceGraph triplestore.RDFGraph,
	mappingGraph triplestore.RDFGraph) {

	var graphSetForMergingMapping []triplestore.RDFGraph

	graphSetForMergingMapping = append(
		graphSetForMergingMapping,
		mergedDifferenceGraph,
		mappingGraph)

	mergedDifferenceGraphWithMapping :=
		MergeGraphs(
			graphSetForMergingMapping)

	DrawGraph(
		mergedDifferenceGraphWithMapping,
		"merged_differences_with_mapping_graph.dot")
}

func MergeGraphs(
	tripleGraphs []triplestore.RDFGraph) triplestore.RDFGraph {

	tripleGraph2Triples := tripleGraphs[1].Triples()
	tripleGraph1Triples := tripleGraphs[0].Triples()

	mergedTripleStore := triplestore.NewSource()

	for _, tripleStore1Triple := range tripleGraph1Triples {
		mergedTripleStore.Add(tripleStore1Triple)
	}

	for _, tripleStore2Triple := range tripleGraph2Triples {
		mergedTripleStore.Add(tripleStore2Triple)
	}

	mergedTripleGraph := mergedTripleStore.Snapshot()

	return mergedTripleGraph
}
