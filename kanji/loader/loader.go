package loader

func Load[T any](m map[rune]T) map[rune]struct{} {
	res := make(map[rune]struct{}, len(m))
	for r := range m {
		res[r] = struct{}{}
	}
	return res
}

func Intersection(m1 map[rune]struct{}, m2 map[rune]struct{}) map[rune]struct{} {
	res := make(map[rune]struct{}, len(m1))
	for r := range m1 {
		if _, ok := m2[r]; ok {
			res[r] = struct{}{}
		}
	}
	return res
}
