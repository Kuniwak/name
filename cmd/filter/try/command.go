package try

import (
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/mora"
)

var SubCommand = cli.SubCommand{
	Help: "try a filter",
	Command: func(args []string, procInout cli.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()

		opts, err := ParseOptions(args, procInout.Stdin, procInout.Stderr, strokesMap)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		evalResult, err := eval.Evaluate(opts.FamilyName, opts.GivenName, strokesMap)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to evaluate: %s\n", err.Error())
			return 1
		}

		if opts.Filter(filter.Target{
			GivenName:  opts.GivenName,
			Yomi:       opts.Yomi,
			YomiString: string(opts.Yomi),
			Strokes:    eval.SumStrokes(opts.GivenName, strokesMap),
			Mora:       mora.Count(opts.Yomi),
			EvalResult: evalResult,
		}) {
			return 0
		}

		return 1
	},
}
