package filter

import (
	"fmt"
	"github.com/Kuniwak/name/eval"
)

type Target struct {
	GivenName  []rune
	Yomi       []rune
	YomiString string
	Strokes    byte
	Mora       byte
	EvalResult eval.Result
}

func (t Target) String() string {
	return fmt.Sprintf("Target{GivenName: %s, Yomi: %s, Strokes: %d, Mora: %d, EvalResult: %s}", string(t.GivenName), string(t.Yomi), t.Strokes, t.Mora, t.EvalResult.String())
}
