package eval

import (
	"fmt"
)

type Rank byte

const (
	DaiDaiKichi Rank = 4
	DaiKichi    Rank = 3
	Kichi       Rank = 2
	Kyo         Rank = 1
	DaiKyo      Rank = 0
)

func (r Rank) String() string {
	switch r {
	case DaiDaiKichi:
		return "大大吉"
	case DaiKichi:
		return "大吉"
	case Kichi:
		return "吉"
	case Kyo:
		return "凶"
	case DaiKyo:
		return "大凶"
	}
	panic("unknown rank")
}

func StrokesToRank(strokes byte) (Rank, error) {
	switch strokes {
	case 15, 24, 31:
		return DaiDaiKichi, nil
	case 1, 3, 5, 6, 11, 13, 16, 21, 23, 29, 32, 33, 35, 37, 39:
		return DaiKichi, nil
	case 7, 8, 17, 18, 25, 26, 38:
		return Kichi, nil
	case 14, 22, 27, 28, 30:
		return Kyo, nil
	case 2, 4, 9, 10, 12, 19, 20, 34, 36, 40:
		return DaiKyo, nil
	}
	return 0, fmt.Errorf("too large strokes: %d", strokes)
}

type Result struct {
	Tenkaku Rank
	Jinkaku Rank
	Chikaku Rank
	Gaikaku Rank
	Sokaku  Rank
}

func (r Result) Total() byte {
	return byte(r.Tenkaku) + byte(r.Jinkaku) + byte(r.Chikaku) + byte(r.Gaikaku) + byte(r.Sokaku)
}

func (r Result) String() string {
	return fmt.Sprintf("Result{Tenkaku: %s, Jinkaku: %s, Chikaku: %s, Gaikaku: %s, Sokaku: %s}", r.Tenkaku.String(), r.Jinkaku.String(), r.Chikaku.String(), r.Gaikaku.String(), r.Sokaku.String())
}

func Evaluate(familyName, givenName []rune, strokeMap map[rune]byte) (Result, error) {
	tenkaku, err := StrokesToRank(Tenkaku(familyName, strokeMap))
	if err != nil {
		return Result{}, err
	}

	jinkaku, err := StrokesToRank(Jinkaku(familyName, givenName, strokeMap))
	if err != nil {
		return Result{}, err
	}

	chikaku, err := StrokesToRank(Chikaku(givenName, strokeMap))
	if err != nil {
		return Result{}, err
	}

	gaikaku, err := StrokesToRank(Gaikaku(familyName, givenName, strokeMap))
	if err != nil {
		return Result{}, err
	}

	sokaku, err := StrokesToRank(Sokaku(familyName, givenName, strokeMap))
	if err != nil {
		return Result{}, err
	}

	return Result{
		Tenkaku: tenkaku,
		Jinkaku: jinkaku,
		Chikaku: chikaku,
		Gaikaku: gaikaku,
		Sokaku:  sokaku,
	}, nil
}

func Strokes(r rune, strokeMap map[rune]byte) byte {
	n, ok := strokeMap[r]
	if !ok {
		panic(fmt.Sprintf("unknown kanji: %c", r))
	}
	return n
}

func SumStrokes(rs []rune, strokeMap map[rune]byte) byte {
	var sum byte = 0
	for _, r := range rs {
		sum += Strokes(r, strokeMap)
	}
	return sum
}

func Tenkaku(familyName []rune, strokeMap map[rune]byte) byte {
	return SumStrokes(familyName, strokeMap)
}

func Jinkaku(familyName, givenName []rune, strokeMap map[rune]byte) byte {
	c1 := familyName[len(familyName)-1]
	c2 := givenName[0]
	return Strokes(c1, strokeMap) + Strokes(c2, strokeMap)
}

func Chikaku(givenName []rune, strokeMap map[rune]byte) byte {
	return SumStrokes(givenName, strokeMap)
}

func Gaikaku(familyName, givenName []rune, strokeMap map[rune]byte) byte {
	c1 := familyName[0]
	c2 := givenName[len(givenName)-1]
	var n1 byte
	if len(familyName) == 1 {
		n1 = Strokes(c1, strokeMap) + 1
	} else {
		n1 = Strokes(c1, strokeMap)
	}
	var n2 byte
	if len(givenName) == 1 {
		n2 = Strokes(c2, strokeMap) + 1
	} else {
		n2 = Strokes(c2, strokeMap)
	}
	return n1 + n2
}

func Sokaku(familyName, givenName []rune, strokeMap map[rune]byte) byte {
	return SumStrokes(familyName, strokeMap) + SumStrokes(givenName, strokeMap)
}
