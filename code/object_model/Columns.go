package object_model

import "database/sql"

type Columns struct {
	IsNullable          bool
	IsTemporal          bool
	IsNullablePrimitive bool
	IsNullableTemporal  bool
}

func (c Columns) HasTrue() bool {

	hasTrue :=
		c.IsNullable ||
			c.IsTemporal ||
			c.IsNullableTemporal ||
			c.IsNullablePrimitive

	return hasTrue
}

// Column stores information about a column
type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	Name                   string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	DefaultValue           sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	ColumnKey              string         `db:"column_key"`      // mysql specific
	Extra                  string         `db:"extra"`           // mysql specific
	ConstraintName         sql.NullString `db:"constraint_name"` // pg specific
	ConstraintType         sql.NullString `db:"constraint_type"` // pg specific
}
