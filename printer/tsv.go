package printer

import (
	"fmt"
	"github.com/Kuniwak/name/filter"
	"io"
)

var TSVHeaders = []string{
	"評点", "画数", "名前", "読み", "性別", "天格", "地格", "人格", "外格", "総格",
}

var tab = []byte("\t")
var newline = []byte("\n")

func PrintTSVHeader(w io.Writer) {
	for i, header := range TSVHeaders {
		if i > 0 {
			_, _ = w.Write(tab)
		}
		_, _ = io.WriteString(w, header)
	}
	_, _ = w.Write(newline)
}

func PrintTSVRow(w io.Writer, d filter.Target) {
	_, _ = fmt.Fprintf(w, "%d", d.EvalResult.Total())
	_, _ = w.Write(tab)
	_, _ = fmt.Fprintf(w, "%d", d.Strokes)
	_, _ = w.Write(tab)
	for _, r := range d.Kanji {
		_, _ = fmt.Fprintf(w, "%c", r)
	}
	_, _ = w.Write(tab)
	for _, r := range d.Yomi {
		_, _ = fmt.Fprintf(w, "%c", r)
	}
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.Sex.String())
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.EvalResult.Tenkaku.String())
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.EvalResult.Chikaku.String())
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.EvalResult.Jinkaku.String())
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.EvalResult.Gaikaku.String())
	_, _ = w.Write(tab)
	_, _ = io.WriteString(w, d.EvalResult.Sokaku.String())
	_, _ = w.Write(newline)
}

func NewTSVPrinter(w io.Writer) Func {
	return func(ch <-chan filter.Target) {
		PrintTSVHeader(w)
		for d := range ch {
			PrintTSVRow(w, d)
		}
	}
}
