package filter

import "github.com/Kuniwak/name/sex"

type SexFunc func(s sex.Sex) bool

func Sex(f SexFunc) Func {
	return func(t Target) bool {
		return f(t.Sex)
	}
}

func Asexual(s sex.Sex) bool {
	switch s {
	case sex.Asexual:
		return true
	default:
		return false
	}
}

func Female(s sex.Sex) bool {
	switch s {
	case sex.Female, sex.Asexual:
		return true
	default:
		return false
	}
}

func Male(s sex.Sex) bool {
	switch s {
	case sex.Male, sex.Asexual:
		return true
	default:
		return false
	}
}
