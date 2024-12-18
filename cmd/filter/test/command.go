package test

import (
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/kanji/loader"
	"github.com/Kuniwak/name/mora"
	"github.com/Kuniwak/name/sex"
	"github.com/Kuniwak/name/strokes"
)

var SubCommand = cli.SubCommand{
	Help: "test a filter",
	Command: func(args []string, procInout cli.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()
		yomiMap := kanji.LoadYomi()
		cm := loader.Intersection(loader.Load(strokesMap), loader.Load(yomiMap))

		opts, err := ParseOptions(args, procInout.Stdin, procInout.Stderr, cm)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		strokesFunc := strokes.ByMap(kanji.LoadStrokes())

		evalResult, err := eval.Evaluate(opts.FamilyName, opts.GivenName, strokesFunc)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to evaluate: %s\n", err.Error())
			return 1
		}

		s, err := strokes.Sum(opts.GivenName, strokesFunc)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to sum strokes: %s\n", err.Error())
			return 1
		}

		yomiString := string(opts.Yomi)

		sexFunc := sex.ByNameLists(sex.LoadMaleNames(), sex.LoadFemaleNames())
		if opts.Filter(filter.Target{
			Kanji:      opts.GivenName,
			Yomi:       opts.Yomi,
			YomiString: yomiString,
			Strokes:    s,
			Sex:        sexFunc(yomiString),
			Mora:       mora.Count(opts.Yomi),
			EvalResult: evalResult,
		}) {
			return 0
		}

		return 1
	},
}
