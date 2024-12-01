package gen

import (
	"github.com/Kuniwak/name/config"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/sliceutil"
)

func NewFullSpaceGenerator(strokes map[rune]byte, yomiMap map[rune][][]rune) GenerateFunc {
	return func(familyName []rune, ch chan<- Generated) {
		m := config.MaxStrokes - eval.SumStrokes(familyName, strokes)

		for r1, stroke1 := range strokes {
			if stroke1 > m {
				continue
			}

			emit(ch, []rune{r1}, yomiMap)

			for r2, stroke2 := range strokes {
				if stroke1+stroke2 > m {
					continue
				}

				emit(ch, []rune{r1, r2}, yomiMap)

				for r3, stroke3 := range strokes {
					if stroke1+stroke2+stroke3 > m {
						continue
					}

					emit(ch, []rune{r1, r2, r3}, yomiMap)
				}
			}
		}

		close(ch)
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
