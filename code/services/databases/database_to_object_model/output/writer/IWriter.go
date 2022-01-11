package writer

// Writer represents an interface to write the produced struct content.
type Writer interface {
	Write(tableName string,
		content string) error
}
