package object_model

import (
	"fmt"
	//"github.com/OntoLedgy/storage_interop_services/code/services/database_transactions_service/internal"
)

type LimitOptions struct {
	Offset int
	Count  int
}

func (limitOptions *LimitOptions) Limit(
	offset, count int) SelectOptions {

	return &LimitOptions{
		offset,
		count}
}

func (limitOptions *LimitOptions) Wrap(
	query string,
	params []interface{}) (
	string,
	[]interface{}) {

	query = fmt.Sprintf("SELECT a.* FROM (%s) a LIMIT $1 OFFSET $2", query)

	params = append(
		params,
		limitOptions.Count)

	params = append(
		params,
		limitOptions.Offset)

	return query, params
}
