package dataframe_helpers

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/dataframes"
	"github.com/go-gota/gota/series"
	"github.com/tobgu/qframe/types"
	"math/rand"
)

//from pandas import DataFrame
//from nf_common_source.code.services.identification_services.uuid_service.uuid_helpers.uuid_factory import \
//create_new_uuid
//
//
//def uuidify_dataframe(
func UuidifyDataframe(
	dataframe dataframes.DataFrames,
	uuidColumnName string) *dataframes.DataFrames {

	uuidifedDataFrame := &dataframes.DataFrames{
		DataFrame: dataframe.DataFrame.Copy()}

	//TODO how to add lambda function in go https://faun.pub/can-we-write-lambda-functions-in-golang-519d44712235

	uuidifedDataFrame.Mutate(
		series.New(rand.Int(),
			types.String,
			uuidColumnName))

	//dataframe: DataFrame,
	//uuid_column_name: str) \
	//-> DataFrame:
	//uuidified_dataframe = \
	//dataframe.copy()
	//
	//uuidified_dataframe[uuid_column_name] = \
	//uuidified_dataframe.apply(
	//lambda row: create_new_uuid(),
	//axis=1)
	//
	//dataframe_without_new_uuid_column = \
	//uuidified_dataframe.pop(
	//uuid_column_name)
	//
	//uuidified_dataframe.insert(
	//0,
	//uuid_column_name,
	//dataframe_without_new_uuid_column)
	//
	//return \
	//uuidified_datafram

	return uuidifedDataFrame
}
