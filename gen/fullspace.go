package gen

import (
	"github.com/Kuniwak/name/config"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/sliceutil"
)

func NewFullSpaceGenerator(strokes map[rune]byte, yomiMap map[rune][][]rune) GenerateFunc {
	return func(familyName []rune, opts Options, ch chan<- Generated) {
		m := config.MaxStrokes - eval.SumStrokes(familyName, strokes)
		fullSpaceGenerator([]rune{}, 0, m, opts, strokes, yomiMap, ch)
		close(ch)
	}
}

func fullSpaceGenerator(current []rune, currentStrokes, maxStrokes byte, opts Options, strokesMap map[rune]byte, yomiMap map[rune][][]rune, ch chan<- Generated) {
	if len(current) >= opts.MaxLength {
		return
	}

	for r, stroke := range strokesMap {
		if currentStrokes+stroke > maxStrokes {
			continue
		}

		newCurrent := append(append([]rune{}, current...), r)
		if len(newCurrent) >= opts.MinLength {
			emit(ch, newCurrent, yomiMap)
		}

		fullSpaceGenerator(newCurrent, currentStrokes+stroke, maxStrokes, opts, strokesMap, yomiMap, ch)
	}
}

func emit(ch chan<- Generated, givenName []rune, yomiMap map[rune][][]rune) {
	rsss := make([][][]rune, len(givenName))

	for i, r := range givenName {
		yomis := yomiMap[r]
		rsss[i] = yomis
	}

	c := sliceutil.Cartesian(rsss)
	yomis := make([][]rune, len(c))
	for i, rs := range c {
		yomis[i] = sliceutil.Flatten(rs)
	}

	for _, yomi := range yomis {
		ch <- Generated{
			GivenName:  givenName,
			Yomi:       yomi,
			YomiString: string(yomi),
		}
	}
}
