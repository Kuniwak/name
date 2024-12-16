package filter

import (
	"fmt"
	"github.com/Kuniwak/name/eval"
)

type Target struct {
	Kanji      []rune
	Yomi       []rune
	YomiString string
	Strokes    byte
	Mora       byte
	EvalResult eval.Result
}

func (t Target) String() string {
	return fmt.Sprintf("Target{Kanji: %s, Yomi: %s, Strokes: %d, Mora: %d, EvalResult: %s}", string(t.Kanji), string(t.Yomi), t.Strokes, t.Mora, t.EvalResult.String())
}
