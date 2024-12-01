package cmd

import (
	"fmt"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/kanji"
	"github.com/Kuniwak/name/printer"
	"github.com/Kuniwak/name/search"
	"io"
	"runtime"
	"sync"
)

func MainCmd(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) byte {
	strokesMap := kanji.LoadStrokes()
	yomiMap := kanji.LoadYomi()

	opts, usage, err := ParseOptions(args, stdin, stderr, strokesMap, yomiMap)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "failed to parse options: %s\n", err.Error())
		return 1
	}

	if opts.Help {
		usage()
		return 0
	}

	Main(opts.FamilyName, opts.GeneratorFunc, opts.Filter, strokesMap, yomiMap, printer.NewTSVPrinter(stdout))
	return 0
}

func Main(familyName []rune, genFunc gen.GenerateFunc, filterFunc filter.Func, strokesMap map[rune]byte, yomiMap map[rune][][]rune, printFunc printer.Func) {
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
