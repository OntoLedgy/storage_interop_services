package graph_processors

import (
	"fmt"
	"github.com/wallix/triplestore"
)

func AnalyseGraphsPopulation(
	inputGraphs []triplestore.RDFGraph) {

	for index, graph := range inputGraphs {

		AnalyseGraphPopulation(graph, index)

	}

}

func AnalyseGraphPopulation(
	graph triplestore.RDFGraph,
	index int) {

	numberOfTriples := len(graph.Triples())

	triples := graph.Triples()

	countOfObjects :=
		countTripleObjects(
			triples)

	countOfSubjects :=
		countTripleSubjects(
			triples)

	countOfPredicates :=
		countTriplePredicates(
			triples)

	fmt.Printf(
		"graph number: %v\nnumber of triples: %v\nnumber of subjects :%v\nnumber of predicates :%v\nnumber of objects :%v\n",
		index,
		numberOfTriples,
		countOfSubjects,
		countOfPredicates,
		countOfObjects)

}

func countTripleObjects(
	triples []triplestore.Triple) int {

	uniqueObjectsMap :=
		map[triplestore.Object]bool{}

	var uniqueObjectsList []triplestore.Object

	//TODO - Generalise this into a set operation

	for _, triple := range triples {

		if !uniqueObjectsMap[triple.Object()] {
			uniqueObjectsMap[triple.Object()] = true
			uniqueObjectsList = append(uniqueObjectsList, triple.Object())
		}

	}
	return len(uniqueObjectsList)
}

func countTripleSubjects(
	triples []triplestore.Triple) int {

	uniqueObjectsMap := map[string]bool{}

	var uniqueObjectsList []string

	for _, triple := range triples {

		if !uniqueObjectsMap[triple.Subject()] {

			uniqueObjectsMap[triple.Subject()] = true

			uniqueObjectsList = append(uniqueObjectsList, triple.Subject())
		}

	}
	return len(uniqueObjectsList)
}

func countTriplePredicates(
	triples []triplestore.Triple) int {

	uniqueObjectsMap := map[string]bool{}

	var uniqueObjectsList []string

	for _, triple := range triples {

		if !uniqueObjectsMap[triple.Predicate()] {
			uniqueObjectsMap[triple.Predicate()] = true
			uniqueObjectsList = append(uniqueObjectsList, triple.Predicate())
		}

	}
	return len(uniqueObjectsList)
}
