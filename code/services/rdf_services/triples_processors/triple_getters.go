package triples_processors

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
	"github.com/wallix/triplestore"
)

func GetTriplesDatasets(
	file1Name,
	file2Name,
	delimiterString string) []triplestore.Source {

	var sourceTriples []triplestore.Source

	//TODO make this an array of array of slices
	triplesDataset1 :=
		csv.ReadCsvToSlice(
			file1Name,
			delimiterString)

	triplesDataset2 :=
		csv.ReadCsvToSlice(
			file2Name,
			delimiterString)

	tripleStore1 :=
		getTripleDataset(
			triplesDataset1)

	tripleStore2 :=
		getTripleDataset(
			triplesDataset2)

	sourceTriples = append(
		sourceTriples,
		tripleStore1,
		tripleStore2)

	return sourceTriples
}

func getTripleDataset(
	triplesDataset [][]string) triplestore.Source {

	tripleStore :=
		triplestore.
			NewSource()

	for _, tripleData := range triplesDataset {

		walliTriple :=
			triplestore.
				SubjPred(
					tripleData[0],
					tripleData[1]).
				Resource(
					tripleData[2])

		tripleStore.
			Add(
				walliTriple)

	}

	return tripleStore
}
