package configurations

type ColumnInfo struct {
	IsNullable          bool
	IsTemporal          bool
	IsNullablePrimitive bool
	IsNullableTemporal  bool
}

func (c ColumnInfo) HasTrue() bool {
	return c.IsNullable || c.IsTemporal || c.IsNullableTemporal || c.IsNullablePrimitive
}
