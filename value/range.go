package value

type ValueRange[T comparable] struct {
	Field        string
	From         T
	To           T
	IncludeEmpty bool
}
