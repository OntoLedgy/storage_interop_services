package dataframes

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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

func (dataframe *DataFrames) FillNa(fillString string) {

	//TODO - make this an explicit function - wrap around series? or review alternative implemtations.
	var replaceSeriesNullWithDefaultNullValue = series.MapFunction(func(element series.Element) series.Element {

		if element.IsNA() || len(element.String()) == 0 {
			switch element.Type() {
			case series.String:
				element.Set(fillString)
				return element
			case series.Int:
				element.Set(0)
				return element
			case series.Float:
				element.Set(0)
				return element
			case series.Bool:
				element.Set(false)
				return element
			}
		}
		return element
	})

	dataframe_columns :=
		dataframe.Columns()

	for _, dataframe_column := range dataframe_columns {

		newSeries := dataframe.DataFrame.Col(dataframe_column).Map(replaceSeriesNullWithDefaultNullValue)

		dataframe.DataFrame = dataframe.Mutate(newSeries)
		fmt.Println(newSeries)
	}

}
