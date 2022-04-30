package ptr

// New returns the pointer of the given value.
func New[T any](v T) *T {
	return &v
}
