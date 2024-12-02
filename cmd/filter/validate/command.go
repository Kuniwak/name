package validate

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/filter"
	"io"
)

var SubCommand = cli.SubCommand{
	Help: "validate a filter",
	Command: func(args []string, procInout cli.ProcInout) byte {
		flags := flag.NewFlagSet("validate", flag.ContinueOnError)
		flags.SetOutput(procInout.Stderr)

		flags.Usage = func() {
			_, _ = procInout.Stderr.Write([]byte("Usage: name filter validate\n"))
			_, _ = fmt.Fprintf(procInout.Stderr, `
EXAMPLES
	$ name filter validate < valid-filter.json
	$ echo $?
	0

	$ name filter validate < invalid-filter.json
	$ echo $?
	1
`)

		}

		if err := flags.Parse(args); err != nil {
			if errors.Is(err, flag.ErrHelp) {
				return 0
			}
		}

		bs, err := io.ReadAll(procInout.Stdin)
		if err != nil {
			_, _ = fmt.Fprintf(procInout.Stderr, "failed to read from stdin: %s\n", err.Error())
			return 1
		}

		var f filter.Data
		if err := json.Unmarshal(bs, &f); err != nil {
			return 1
		}

		d, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			return 1
		}
	
		_, _ = procInout.Stdout.Write(d)
		return 0
	},
}
