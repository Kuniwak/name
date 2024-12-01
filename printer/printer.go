package printer

import (
	"fmt"
	"github.com/Kuniwak/name/filter"
	"io"
)

type Func func(<-chan filter.Target)

func NewTSVPrinter(w io.Writer) Func {
	return func(ch <-chan filter.Target) {
		_, _ = io.WriteString(w, "評点\t画数\t名前\t読み\t天格\t地格\t人格\t外格\t総格\n")
		for d := range ch {
			_, _ = fmt.Fprintf(w, "%d\t", d.EvalResult.Total())
			_, _ = fmt.Fprintf(w, "%d\t", d.Strokes)
			for _, r := range d.GivenName {
				_, _ = fmt.Fprintf(w, "%c", r)
			}
			_, _ = w.Write([]byte("\t"))
			for _, r := range d.Yomi {
				_, _ = fmt.Fprintf(w, "%c", r)
			}
			_, _ = w.Write([]byte("\t"))
			_, _ = io.WriteString(w, d.EvalResult.Tenkaku.String())
			_, _ = w.Write([]byte("\t"))
			_, _ = io.WriteString(w, d.EvalResult.Chikaku.String())
			_, _ = w.Write([]byte("\t"))
			_, _ = io.WriteString(w, d.EvalResult.Jinkaku.String())
			_, _ = w.Write([]byte("\t"))
			_, _ = io.WriteString(w, d.EvalResult.Gaikaku.String())
			_, _ = w.Write([]byte("\t"))
			_, _ = io.WriteString(w, d.EvalResult.Sokaku.String())
			_, _ = w.Write([]byte("\n"))
		}
	}
}
