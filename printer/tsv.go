package printer

import (
	"fmt"
	"github.com/Kuniwak/name/filter"
	"io"
)

func PrintTSVHeader(w io.Writer) {
	_, _ = io.WriteString(w, "評点\t画数\t名前\t読み\t天格\t地格\t人格\t外格\t総格\n")
}

var tab = []byte("\t")
var newline = []byte("\n")

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
