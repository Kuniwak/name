package yomi

import (
	"fmt"
	"github.com/Kuniwak/name/sliceutil"
	"github.com/Kuniwak/name/yomi/namelti"
	"github.com/shogo82148/go-mecab"
)

type Func func(rs []rune) ([]Result, error)

type Result struct {
	Runes  []rune
	String string
}

func ByConstant(yomi []string) Func {
	return func(rs []rune) ([]Result, error) {
		results := make([]Result, len(yomi))
		for i, y := range yomi {
			results[i] = Result{
				Runes:  []rune(y),
				String: y,
			}
		}
		return results, nil
	}
}

func Fallback(f, fallback Func) Func {
	return func(rs []rune) ([]Result, error) {
		yomis, err := f(rs)
		if err != nil {
			return nil, err
		}

		if len(yomis) > 0 {
			return yomis, nil
		}
		return fallback(rs)
	}
}

func ByCartesian(yomiDict map[rune][][]rune) Func {
	return func(rs []rune) ([]Result, error) {
		yomis := make([][][]rune, len(rs))
		for i, r := range rs {
			yomis[i] = yomiDict[r]
		}

		cartesian := sliceutil.Cartesian(yomis)
		results := make([]Result, len(cartesian))
		for i, yomiParts := range cartesian {
			yomi := sliceutil.Flatten(yomiParts)
			results[i] = Result{
				Runes:  yomi,
				String: string(yomi),
			}
		}

		return results, nil
	}
}

func ByMeCab(m mecab.MeCab, nBest int) (Func, error) {
	return func(rs []rune) ([]Result, error) {
		t, err := namelti.NewTranscripter(m)
		if err != nil {
			return nil, fmt.Errorf("failed to create a transcripter: %w", err)
		}

		yomis, err := t.Transcript(string(rs), nBest)
		if err != nil {
			return nil, fmt.Errorf("failed to transcript %q: %w", string(rs), err)
		}

		results := make([]Result, len(yomis))
		for i, yomi := range yomis {
			if yomi == "" {
				continue
			}
			results[i] = Result{
				Runes:  []rune(yomi),
				String: yomi,
			}
		}
		return results, nil
	}, nil
}
