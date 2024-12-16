package search

import (
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/mora"
)

func Search(
	familyName []rune,
	in <-chan gen.Generated,
	out chan<- filter.Target,
	filterFunc filter.Func,
	strokesMap map[rune]byte,
) {
	for generated := range in {
		res, err := eval.Evaluate(familyName, generated.GivenName, strokesMap)
		if err != nil {
			continue
		}

		target := filter.Target{
			Kanji:      generated.GivenName,
			Yomi:       generated.Yomi,
			YomiString: generated.YomiString,
			Strokes:    eval.SumStrokes(generated.GivenName, strokesMap),
			Mora:       mora.Count(generated.Yomi),
			EvalResult: res,
		}
		if filterFunc(target) {
			out <- target
		}
	}

	close(out)
}
