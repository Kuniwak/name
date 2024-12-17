package search

import (
	"fmt"
	cli2 "github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/kanji/loader"
	"github.com/Kuniwak/name/printer"
	"github.com/Kuniwak/name/search"
	"github.com/Kuniwak/name/sex"
	"github.com/Kuniwak/name/strokes"
	"github.com/Kuniwak/name/yomi"
	"golang.org/x/sync/errgroup"
	"runtime"
)

var SubCommand = cli2.SubCommand{
	Help: "search for given names",
	Command: func(args []string, procInout cli2.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()
		strokesFunc := strokes.ByMap(strokesMap)
		yomiMap := kanji.LoadYomi()
		yomiFunc := yomi.ByCartesian(yomiMap)
		cm := loader.Intersection2(loader.Load(strokesMap), loader.Load(yomiMap))

		opts, err := ParseOptions(args, procInout.Stdin, procInout.Stderr, cm, strokesFunc, yomiFunc)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		genOpts := gen.Options{
			MinLength: opts.MinLength,
			MaxLength: opts.MaxLength,
		}

		if err := Main(
			opts.FamilyName,
			opts.GeneratorFunc,
			genOpts,
			opts.Filter,
			strokesFunc,
			printer.NewTSVPrinter(procInout.Stdout),
			sex.ByNameLists(sex.LoadMaleNames(), sex.LoadFemaleNames()),
		); err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to search: %s\n", err.Error())
			return 1
		}
		return 0
	},
}

func Main(familyName []rune, genFunc gen.GenerateFunc, genOpts gen.Options, filterFunc filter.Func, strokesFunc strokes.Func, printFunc printer.Func, sexFunc sex.Func) error {
	candCh := make(chan gen.Generated)
	resCh := make(chan filter.Target)

	eg := new(errgroup.Group)

	eg.Go(func() error {
		return genFunc(familyName, genOpts, candCh)
	})

	eg.Go(func() error {
		printFunc(resCh)
		return nil
	})

	parallelism := runtime.NumCPU() - 2
	if parallelism < 1 {
		parallelism = 1
	}

	if err := search.Parallel(familyName, candCh, resCh, filterFunc, strokesFunc, sexFunc, parallelism); err != nil {
		return err
	}

	return eg.Wait()
}
