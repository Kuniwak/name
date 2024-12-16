package info

import (
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/mora"
	"github.com/Kuniwak/name/printer"
)

var SubCommand = cli.SubCommand{
	Help: "show information about a given name",
	Command: func(args []string, procInout cli.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()

		opts, err := ParseOptions(args, procInout.Stderr, strokesMap)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		printer.PrintTSVHeader(procInout.Stdout)

		evalResult, err := eval.Evaluate(opts.FamilyName, opts.GivenName, strokesMap)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to evaluate: %s\n", err.Error())
			return 1
		}

		target := filter.Target{
			Kanji:      opts.GivenName,
			Yomi:       opts.Yomi,
			YomiString: string(opts.Yomi),
			Strokes:    eval.SumStrokes(opts.GivenName, strokesMap),
			Mora:       mora.Count(opts.Yomi),
			EvalResult: evalResult,
		}
		printer.PrintTSVRow(procInout.Stdout, target)
		return 0
	},
}
