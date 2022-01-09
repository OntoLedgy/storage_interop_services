package data_model

type SelectOptions interface {
	Wrap(
		string,
		[]interface{}) (
		string,
		[]interface{})
}
