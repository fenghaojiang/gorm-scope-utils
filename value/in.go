package value

type ValueIn[T comparable] struct {
	Field  string
	Values []T
}
