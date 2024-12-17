package strokes

import "fmt"

type Func func(rs rune) (byte, error)

func Sum(rs []rune, f Func) (byte, error) {
	var sum byte
	for _, r := range rs {
		strokes, err := f(r)
		if err != nil {
			return 0, err
		}
		sum += strokes
	}
	return sum, nil
}

func Add(r1 rune, r2 rune, f Func) (byte, error) {
	strokes1, err := f(r1)
	if err != nil {
		return 0, err
	}
	strokes2, err := f(r2)
	if err != nil {
		return 0, err
	}
	return strokes1 + strokes2, nil
}

func ByMap(strokesMap map[rune]byte) Func {
	return func(r rune) (byte, error) {
		if strokes, ok := strokesMap[r]; ok {
			return strokes, nil
		} else {
			return 0, fmt.Errorf("strokes not found for %v", r)
		}
	}
}

func ByConstant(strokes byte, err error) Func {
	return func(r rune) (byte, error) {
		return strokes, err
	}
}
