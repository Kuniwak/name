package loader

func Load[T any](m map[rune]T) map[rune]struct{} {
	res := make(map[rune]struct{}, len(m))
	for r := range m {
		res[r] = struct{}{}
	}
	return res
}

func Union(ms ...map[rune]struct{}) map[rune]struct{} {
	l := 0
	for _, m := range ms {
		l += len(m)
	}

	res := make(map[rune]struct{}, l)
	for _, m := range ms {
		for r := range m {
			res[r] = struct{}{}
		}
	}
	return res
}

func Intersection(ms ...map[rune]struct{}) map[rune]struct{} {
	if len(ms) == 0 {
		return nil
	}

	m0 := ms[0]
	return Intersection2(m0, Intersection(ms[1:]...))
}

func Intersection2(m1 map[rune]struct{}, m2 map[rune]struct{}) map[rune]struct{} {
	res := make(map[rune]struct{}, len(m1))
	for r := range m1 {
		if _, ok := m2[r]; ok {
			res[r] = struct{}{}
		}
	}
	return res
}
