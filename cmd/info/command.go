package info

import (
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/kanji/loader"
	"github.com/Kuniwak/name/mora"
	"github.com/Kuniwak/name/printer"
	"github.com/Kuniwak/name/sex"
	"github.com/Kuniwak/name/strokes"
)

var SubCommand = cli.SubCommand{
	Help: "show information about a given name",
	Command: func(args []string, procInout cli.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()
		strokesFunc := strokes.ByMap(strokesMap)
		yomiMap := kanji.LoadYomi()
		cm := loader.Intersection(loader.Load(strokesMap), loader.Load(yomiMap))

		opts, err := ParseOptions(args, procInout.Stderr, cm)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		yomiString := string(opts.Yomi)
		printer.PrintTSVHeader(procInout.Stdout)

		evalResult, err := eval.Evaluate(opts.FamilyName, opts.GivenName, strokesFunc)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to evaluate: %s\n", err.Error())
			return 1
		}

		sexFunc := sex.ByNameLists(sex.LoadMaleNames(), sex.LoadFemaleNames())

		s, err := strokes.Sum(opts.GivenName, strokesFunc)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to sum strokes: %s\n", err.Error())
			return 1
		}

		target := filter.Target{
			Kanji:      opts.GivenName,
			Yomi:       opts.Yomi,
			YomiString: yomiString,
			Strokes:    s,
			Mora:       mora.Count(opts.Yomi),
			Sex:        sexFunc(yomiString),
			EvalResult: evalResult,
		}
		printer.PrintTSVRow(procInout.Stdout, target)
		return 0
	},
}
