package object_model

// Table has a name and a set (slice) of columns
type Table struct {
	Name    string `db:"table_name"`
	Columns []Column
}
