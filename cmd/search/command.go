package search

import (
	"fmt"
	cli2 "github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/printer"
	"github.com/Kuniwak/name/search"
	"runtime"
	"sync"
)

var SubCommand = cli2.SubCommand{
	Help: "search for given names",
	Command: func(args []string, procInout cli2.ProcInout) byte {
		strokesMap := kanji.LoadStrokes()
		yomiMap := kanji.LoadYomi()

		opts, err := ParseOptions(args, procInout.Stdin, procInout.Stderr, strokesMap, yomiMap)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		Main(opts.FamilyName, opts.GeneratorFunc, opts.Filter, strokesMap, printer.NewTSVPrinter(procInout.Stdout))
		return 0
	},
}

func Main(familyName []rune, genFunc gen.GenerateFunc, filterFunc filter.Func, strokesMap map[rune]byte, printFunc printer.Func) {
	candCh := make(chan gen.Generated)
	resCh := make(chan filter.Target)

	go genFunc(familyName, candCh)

	go printFunc(resCh)

	parallelism := runtime.NumCPU() - 2
	if parallelism < 1 {
		parallelism = 1
	}

	var wg sync.WaitGroup
	wg.Add(parallelism)
	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg.Done()
			search.Search(familyName, candCh, resCh, filterFunc, strokesMap)
		}()
	}

	wg.Wait()
	close(resCh)
}
