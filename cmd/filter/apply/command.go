package apply

import (
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/parser"
	"github.com/Kuniwak/name/printer"
	"runtime"
	"sync"
)

var SubCommand = cli.SubCommand{
	Help: "apply a filter to name search results",
	Command: func(args []string, procInout cli.ProcInout) byte {
		opts, err := ParseOptions(args, procInout.Stdin, procInout.Stderr)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to parse options: %s\n", err.Error())
			return 1
		}

		if opts.Help {
			return 0
		}

		if err := Main(opts, printer.NewTSVPrinter(procInout.Stdout)); err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to apply filter: %s\n", err.Error())
			return 1
		}

		return 0
	},
}

func Main(opts Options, printFunc printer.Func) error {
	targets := make(chan filter.Target)
	toPrint := make(chan filter.Target)

	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		printFunc(toPrint)
	}()

	parallelism := runtime.NumCPU() - 2
	if parallelism < 1 {
		parallelism = 1
	}
	var wg2 sync.WaitGroup
	wg2.Add(parallelism)
	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg2.Done()
			for target := range targets {
				if opts.Filter(target) {
					toPrint <- target
				}
			}
		}()
	}

	if err := parser.ParseTSV(opts.Result, targets); err != nil {
		return err
	}

	wg2.Wait()
	close(toPrint)

	wg1.Wait()
	return nil
}
