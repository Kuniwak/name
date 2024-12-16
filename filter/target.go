package filter

import (
	"fmt"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/sex"
)

type Target struct {
	Kanji      []rune
	Yomi       []rune
	YomiString string
	Strokes    byte
	Mora       byte
	Sex        sex.Sex
	EvalResult eval.Result
}

func (t Target) String() string {
	return fmt.Sprintf("%s %s %d %d %s %s", string(t.Kanji), t.YomiString, t.Strokes, t.Mora, t.Sex.String(), t.EvalResult.String())
}
