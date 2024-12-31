package eval

import (
	"fmt"
	"github.com/Kuniwak/name/strokes"
)

type Rank byte

const (
	Unknown     Rank = 255
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
	default:
		return "不明"
	}
}

func StrokesToRank(strokes byte) Rank {
	switch strokes {
	case 15, 24, 31:
		return DaiDaiKichi
	case 1, 3, 5, 6, 11, 13, 16, 21, 23, 29, 32, 33, 35, 37, 39:
		return DaiKichi
	case 7, 8, 17, 18, 25, 26, 38:
		return Kichi
	case 14, 22, 27, 28, 30:
		return Kyo
	case 2, 4, 9, 10, 12, 19, 20, 34, 36, 40:
		return DaiKyo
	}
	return Unknown
}

type Result struct {
	Tenkaku Rank
	Jinkaku Rank
	Chikaku Rank
	Gaikaku Rank
	Sokaku  Rank
}

func (r Result) HasUnknown() bool {
	return r.Tenkaku == Unknown || r.Jinkaku == Unknown || r.Chikaku == Unknown || r.Gaikaku == Unknown || r.Sokaku == Unknown
}

func (r Result) Total() byte {
	if r.HasUnknown() {
		return 0
	}
	return byte(r.Tenkaku) + byte(r.Jinkaku) + byte(r.Chikaku) + byte(r.Gaikaku) + byte(r.Sokaku)
}

func (r Result) String() string {
	return fmt.Sprintf("Result{Tenkaku: %s, Jinkaku: %s, Chikaku: %s, Gaikaku: %s, Sokaku: %s}", r.Tenkaku.String(), r.Jinkaku.String(), r.Chikaku.String(), r.Gaikaku.String(), r.Sokaku.String())
}

func Evaluate(familyName, givenName []rune, strokesFunc strokes.Func) (Result, error) {
	tenkakuStrokes, err := Tenkaku(familyName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	tenkaku := StrokesToRank(tenkakuStrokes)

	jinkakuStrokes, err := Jinkaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	jinkaku := StrokesToRank(jinkakuStrokes)

	chikakuStrokes, err := Chikaku(givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	chikaku := StrokesToRank(chikakuStrokes)

	gaikakuStrokes, err := Gaikaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	gaikaku := StrokesToRank(gaikakuStrokes)

	sokakuStrokes, err := Sokaku(familyName, givenName, strokesFunc)
	if err != nil {
		return Result{}, err
	}
	sokaku := StrokesToRank(sokakuStrokes)

	return Result{
		Tenkaku: tenkaku,
		Jinkaku: jinkaku,
		Chikaku: chikaku,
		Gaikaku: gaikaku,
		Sokaku:  sokaku,
	}, nil
}

func Tenkaku(familyName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Sum(familyName, strokesFunc)
}

func Jinkaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Add(familyName[len(familyName)-1], givenName[0], strokesFunc)
}

func Chikaku(givenName []rune, strokesFunc strokes.Func) (byte, error) {
	return strokes.Sum(givenName, strokesFunc)
}

func Gaikaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	c1, err := strokesFunc(familyName[0])
	if err != nil {
		return 0, err
	}

	c2, err := strokesFunc(givenName[len(givenName)-1])
	if err != nil {
		return 0, err
	}

	var n1 byte
	if len(familyName) == 1 {
		n1 = c1 + 1
	} else {
		n1 = c1
	}
	var n2 byte
	if len(givenName) == 1 {
		n2 = c2 + 1
	} else {
		n2 = c2
	}
	return n1 + n2, nil
}

func Sokaku(familyName, givenName []rune, strokesFunc strokes.Func) (byte, error) {
	c1, err := strokes.Sum(familyName, strokesFunc)
	if err != nil {
		return 0, err
	}
	c2, err := strokes.Sum(givenName, strokesFunc)
	if err != nil {
		return 0, err
	}
	return c1 + c2, nil
}
