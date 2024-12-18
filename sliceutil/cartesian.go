package sliceutil

func Cartesian[T any](input ...[]T) [][]T {
	if len(input) == 0 {
		return [][]T{}
	}

	return cartesian(input, []T{})
}

func cartesian[T any](remaining [][]T, prefix []T) [][]T {
	if len(remaining) == 0 {
		return [][]T{prefix}
	}

	current := remaining[0]
	rest := remaining[1:]

	result := make([][]T, 0)
	for _, val := range current {
		newPrefix := append(append([]T{}, prefix...), val)
		result = append(result, cartesian[T](rest, newPrefix)...)
	}
	return result
}
