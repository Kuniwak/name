package gen

import (
	"github.com/Kuniwak/name/config"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/namelti"
	"github.com/shogo82148/go-mecab"
)

func NewFullSpaceGenerator(strokes map[rune]byte, dictDir string, nBest int) (GenerateFunc, error) {
	tagger, err := mecab.New(map[string]string{"dicdir": dictDir})
	if err != nil {
		return nil, err
	}

	return func(familyName []rune, opts Options, ch chan<- Generated) {
		m := config.MaxStrokes - eval.SumStrokes(familyName, strokes)
		fullSpaceGenerator([]rune{}, 0, m, opts, strokes, tagger, nBest, ch)
		close(ch)
	}, nil
}

func fullSpaceGenerator(current []rune, currentStrokes, maxStrokes byte, opts Options, strokesMap map[rune]byte, tagger mecab.MeCab, nBest int, ch chan<- Generated) {
	if len(current) >= opts.MaxLength {
		return
	}

	for r, stroke := range strokesMap {
		if currentStrokes+stroke > maxStrokes {
			continue
		}

		newCurrent := append(append([]rune{}, current...), r)
		if len(newCurrent) >= opts.MinLength {
			if err := emit(ch, newCurrent, tagger, nBest); err != nil {

			}
		}

		fullSpaceGenerator(newCurrent, currentStrokes+stroke, maxStrokes, opts, strokesMap, tagger, nBest, ch)
	}
}

func emit(ch chan<- Generated, givenName []rune, tagger mecab.MeCab, nBest int) error {
	t, err := namelti.NewTranscripter(tagger)
	if err != nil {
		return err
	}
	yomis, err := t.Transcript(string(givenName), nBest)
	if err != nil {
		return err
	}

	for _, yomi := range yomis {
		ch <- Generated{
			GivenName:  givenName,
			Yomi:       []rune(yomi),
			YomiString: yomi,
		}
	}
	return nil
}
