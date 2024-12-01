package sliceutil

// Flatten returns a new slice that is the concatenation of all the slices in ss.
func Flatten[T any](ss [][]T) []T {
	var c int
	for _, s := range ss {
		c += len(s)
	}

	result := make([]T, 0, c)
	for _, s := range ss {
		result = append(result, s...)
	}
	return result
}
