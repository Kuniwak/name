package search

import (
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/mora"
	"github.com/Kuniwak/name/sex"
	"github.com/Kuniwak/name/strokes"
	"golang.org/x/sync/errgroup"
)

func Parallel(
	familyName []rune,
	in <-chan gen.Generated,
	out chan<- filter.Target,
	filterFunc filter.Func,
	strokesFunc strokes.Func,
	sexFunc sex.Func,
	parallelism int,
) error {
	defer close(out)
	var eg errgroup.Group
	for i := 0; i < parallelism; i++ {
		eg.Go(func() error {
			return Search(familyName, in, out, filterFunc, strokesFunc, sexFunc)
		})
	}
	return eg.Wait()
}

func Search(
	familyName []rune,
	in <-chan gen.Generated,
	out chan<- filter.Target,
	filterFunc filter.Func,
	strokesFunc strokes.Func,
	sexFunc sex.Func,
) error {
	for generated := range in {
		res, err := eval.Evaluate(familyName, generated.GivenName, strokesFunc)
		if err != nil {
			return err
		}

		s, err := strokes.Sum(generated.GivenName, strokesFunc)
		if err != nil {
			return err
		}

		target := filter.Target{
			Kanji:      generated.GivenName,
			Yomi:       generated.Yomi,
			YomiString: generated.YomiString,
			Strokes:    s,
			Mora:       mora.Count(generated.Yomi),
			Sex:        sexFunc(generated.YomiString),
			EvalResult: res,
		}
		if filterFunc(target) {
			out <- target
		}
	}
	return nil
}
