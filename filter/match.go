package filter

type MatchFunc func(target []rune) bool

func MatchExactly(matching []rune) MatchFunc {
	return func(target []rune) bool {
		return matchExactly(target, matching)
	}
}

func matchExactly(target, matching []rune) bool {
	if len(target) != len(matching) {
		return false
	}
	for i, r := range target {
		if r != matching[i] {
			return false
		}
	}
	return true
}

func MatchStartsWith(matching []rune) MatchFunc {
	return func(target []rune) bool {
		return matchStartsWith(target, matching)
	}
}

func matchStartsWith(target, matching []rune) bool {
	if len(target) > len(matching) {
		return false
	}
	for i, r := range matching {
		if r != target[i] {
			return false
		}
	}
	return true
}

func MatchEndsWith(matching []rune) MatchFunc {
	return func(target []rune) bool {
		return matchEndsWith(target, matching)
	}
}

func matchEndsWith(target, matching []rune) bool {
	if len(target) > len(matching) {
		return false
	}
	for i, r := range matching {
		if r != target[len(target)-len(matching)+i] {
			return false
		}
	}
	return true
}

func MatchContains(matching []rune) MatchFunc {
	return func(target []rune) bool {
		return matchContains(target, matching)
	}
}

func matchContains(target, matching []rune) bool {
	if len(matching) > len(target) {
		return false
	}
	for i := 0; i <= len(target)-len(matching); i++ {
		if matchStartsWith(target[i:], target) {
			return true
		}
	}
	return false
}
