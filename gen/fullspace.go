package gen

import (
	"github.com/Kuniwak/name/config"
	"github.com/Kuniwak/name/strokes"
	"github.com/Kuniwak/name/yomi"
)

func NewFullSpaceGenerator(cm map[rune]struct{}, strokesFunc strokes.Func, yomiFunc yomi.Func) (GenerateFunc, error) {
	return func(familyName []rune, opts Options, ch chan<- Generated) error {
		defer close(ch)
		s, err := strokes.Sum(familyName, strokesFunc)
		if err != nil {
			return err
		}

		m := config.MaxStrokes - s
		if err := fullSpaceGenerator([]rune{}, 0, m, opts, cm, strokesFunc, yomiFunc, ch); err != nil {
			return err
		}
		return nil
	}, nil
}

func fullSpaceGenerator(current []rune, currentStrokes, maxStrokes byte, opts Options, cm map[rune]struct{}, strokesFunc strokes.Func, yomiFunc yomi.Func, ch chan<- Generated) error {
	if len(current) >= opts.MaxLength {
		return nil
	}

	for r := range cm {
		s, err := strokesFunc(r)
		if err != nil {
			return err
		}

		if currentStrokes+s > maxStrokes {
			continue
		}

		newCurrent := append(append([]rune{}, current...), r)
		if len(newCurrent) >= opts.MinLength {
			if err := emit(ch, newCurrent, yomiFunc); err != nil {
				return err
			}
		}

		if err := fullSpaceGenerator(newCurrent, currentStrokes+s, maxStrokes, opts, cm, strokesFunc, yomiFunc, ch); err != nil {
			return err
		}
	}

	return nil
}

func emit(ch chan<- Generated, givenName []rune, yomiFunc yomi.Func) error {
	yomis, err := yomiFunc(givenName)
	if err != nil {
		return err
	}

	for _, y := range yomis {
		ch <- Generated{
			GivenName:  givenName,
			Yomi:       y.Runes,
			YomiString: y.String,
		}
	}
	return nil
}
