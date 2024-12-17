package sliceutil

func FromChan[T any](ch <-chan T) []T {
	var xs []T
	for x := range ch {
		xs = append(xs, x)
	}
	return xs
}
