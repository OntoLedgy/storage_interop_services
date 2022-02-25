package dataframes

//import uuid
//import pandas

//from nf_common_source.code.constants.nf_common_global_constants import UUIDS_COLUMN_NAME
//from nf_common_source.code.services.identification_services.uuid_service.uuid_helpers.uuid_factory import create_uuid_from_canonical_format_string

import (
	"github.com/OntoLedgy/ol_common_services/code/constants"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/lists"
	"strings"
)

//PARENT_DATAFRAME_SUFFIX = \
//'_merge_parent'
const PARENT_DATAFRAME_SUFFIX = "_merge_parent"

//def drop_duplicated_parent_columns(
//dataframe: pandas.DataFrame):
func DropDuplicatedParentColumns(
	dataframe *DataFrames) {

	//drop_columns_by_marker(
	DropColumnsByMarker(
		//dataframe,
		dataframe,
		//PARENT_DATAFRAME_SUFFIX)
		PARENT_DATAFRAME_SUFFIX)

}

//def drop_columns_by_marker(
func DropColumnsByMarker(
	//dataframe: pandas.DataFrame,
	dataframe *DataFrames,
	//drop_marker: str):
	dropMarker string) {

	//for column in dataframe.columns.values:
	for _, column := range dataframe.Columns() {
		//if drop_marker in column:
		if strings.Contains(column, dropMarker) {
			dataframe.DataFrame.Drop(
				column)
			//dataframe.drop(
			//column,
			//axis=1,
			//inplace=True)

		}

	}

}

//def add_fk_uuids(
func AddForeignKeyUuids(
	//base_dataframe: pandas.DataFrame,
	baseDataFrame DataFrames,
	//parent_register_dataframe: pandas.DataFrame,
	parentRegisterDataFrame DataFrames,
	//base_column_foreign_key: str,
	baseColumnForeignKey string,
	//parent_column_key: str,
	ParentColumnKey string,
	//fk_uuid_column: str,
	ForeignKeyUuidColumn string,
	//remove_fk_after_uuidification):
	removeForeignKeyAferUuidfication bool) *DataFrames {

	//base_dataframe_columns = \
	//list(
	//base_dataframe.columns)

	baseDataFrameColumns :=
		lists.Lists{}

	baseDataFrameColumns.AddStringSlice(
		baseDataFrame.ColumnNames())

	//if remove_fk_after_uuidification:
	if removeForeignKeyAferUuidfication {

		//base_dataframe_columns.remove(
		//base_column_foreign_key)
		baseDataFrameColumns.RemoveElementString(baseColumnForeignKey)
	}

	//merged_dataframe = \
	mergedDataFrame := DataFrames{
		//base_dataframe.merge(
		DataFrame: baseDataFrame.LeftJoin(
			//right=parent_register_dataframe,
			parentRegisterDataFrame.DataFrame,
			//left_on=base_column_foreign_key,
			baseColumnForeignKey,
			//how='left',
			//right_on=parent_column_key,
			//suffixes=['', PARENT_DATAFRAME_SUFFIX]) //TODO wrap this in the dataframe class
			ParentColumnKey)}

	//base_dataframe_columns.append(
	baseDataFrameColumns.PushBack(
		//fk_uuid_column)
		ForeignKeyUuidColumn)

	//if UUIDS_COLUMN_NAME in base_dataframe_columns:
	if baseDataFrameColumns.Contains(constants.UUIDS_COLUMN_NAME) {
		//merged_dataframe.rename(
		//columns={UUIDS_COLUMN_NAME + PARENT_DATAFRAME_SUFFIX: fk_uuid_column},
		//inplace=True)
		mergedDataFrame.Rename(constants.UUIDS_COLUMN_NAME+PARENT_DATAFRAME_SUFFIX, ForeignKeyUuidColumn)
		//else:
	} else {
		//merged_dataframe.rename(
		mergedDataFrame.Rename(
			//columns={UUIDS_COLUMN_NAME: fk_uuid_column},
			constants.UUIDS_COLUMN_NAME, ForeignKeyUuidColumn)
		//inplace=True)
	}

	//fk_uuidified_dataframe = \
	//merged_dataframe.loc[:, base_dataframe_columns]

	foreignKeyUuidifiedDataFrame := &DataFrames{
		DataFrame: mergedDataFrame.DataFrame.Select(baseDataFrameColumns)}

	//return \
	//fk_uuidified_dataframe
	return foreignKeyUuidifiedDataFrame

}

//def add_parent_table_columns(
//base_dataframe: pandas.DataFrame,
//parent_register_dataframe: pandas.DataFrame,
//base_column_foreign_keys: list,
//parent_column_keys: list):
//
//merged_dataframe = \
//base_dataframe.merge(
//right=parent_register_dataframe,
//how="left",
//left_on=base_column_foreign_keys,
//right_on=parent_column_keys,
//suffixes=['', PARENT_DATAFRAME_SUFFIX])
//
//drop_duplicated_parent_columns(
//merged_dataframe)
//
//return \
//merged_dataframe

//def add_type_column_to_dataframe(
//dataframe: pandas.DataFrame,
//col_name: str,
//col_position,
//default_value: uuid):
//
//dataframe.insert(
//col_position,
//col_name,
//default_value)
//
//return \
//dataframe

//def move_uuid_col_to_front(
//dataframe: pandas.DataFrame,
//uuid_column: str):
//columns = \
//list(
//dataframe)
//
//columns.insert(
//0,
//columns.pop(columns.index(uuid_column)))
//
//dataframe = \
//dataframe.loc[:, columns]
//
//return \
//dataframe

//def deduplicate(
//dataframe: pandas.DataFrame,
//columns: list):
//stringified_columns = \
//list()
//
//for column in columns:
//stringified_column = \
//'stringified_'+column
//
//dataframe[stringified_column] = \
//dataframe[column].astype(str)
//
//stringified_columns.append(
//stringified_column)
//
//dataframe.drop_duplicates(
//subset=stringified_columns,
//inplace=True)
//
//dataframe.drop(
//columns=stringified_columns,
//inplace=True)

//def stringify_uuid_columns(
//dataframe: pandas.DataFrame,
//columns: list):
//for column in columns:
//dataframe[column] = \
//dataframe[column].astype(str)

//def unstringify_uuid_columns(
//dataframe: pandas.DataFrame,
//columns: list):
//
//for column in columns:
//dataframe[column] = \
//dataframe[column].apply(
//lambda x: create_uuid_from_canonical_format_string(x))
