package dataframes

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/tobgu/qframe"
)

type DataFrames struct {
	*qframe.QFrame
	dataframe.DataFrame
}

func (dataframe DataFrames) Columns() []string {

	values := dataframe.DataFrame.Names()

	return values
}
