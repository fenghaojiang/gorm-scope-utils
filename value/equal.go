package value

type ValueEqual[T comparable] struct {
	Field        string
	Value        T
	IncludeEmpty bool
}
