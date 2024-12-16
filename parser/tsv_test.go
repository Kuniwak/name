package parser

import (
	"bytes"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/printer"
	"github.com/Kuniwak/name/sex"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestParseTSV(t *testing.T) {
	target := filter.Target{
		Kanji:      []rune("太郎"),
		Yomi:       []rune("タロウ"),
		YomiString: "タロウ",
		Strokes:    13,
		Mora:       3,
		Sex:        sex.Male,
		EvalResult: eval.Result{
			Tenkaku: eval.Kichi,
			Jinkaku: eval.DaiKyo,
			Chikaku: eval.DaiKichi,
			Gaikaku: eval.DaiKyo,
			Sokaku:  eval.DaiKichi,
		},
	}
	buf := &bytes.Buffer{}
	printer.PrintTSVHeader(buf)
	printer.PrintTSVRow(buf, target)

	ch := make(chan filter.Target)
	go func() {
		err := ParseTSV(buf, ch)
		if err != nil {
			t.Error(err.Error())
			return
		}
	}()

	actual := exhaust(ch)
	expected := []filter.Target{target}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(cmp.Diff(expected, actual))
	}
}

func exhaust(ch chan filter.Target) []filter.Target {
	var res []filter.Target
	for target := range ch {
		res = append(res, target)
	}
	return res
}
