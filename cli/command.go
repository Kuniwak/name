package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Command func(args []string, procInout ProcInout) byte

func (c Command) Run() {
	exitStatus := c(os.Args[1:], DefaultProcInout())
	os.Exit(int(exitStatus))
}

type SubCommand struct {
	Help    string
	Command Command
}

func Usage(w io.Writer, name string, d map[string]SubCommand) {
	_, _ = fmt.Fprintf(w, "Usage: %s [subcommand] [options]\n\n", name)
	_, _ = fmt.Fprintf(w, "SUBCOMMANDS\n")
	for k, v := range d {
		_, _ = fmt.Fprintf(w, "  %s    %s\n", k, v.Help)
	}
}

func CommandWithSubCommands(name string, subCommands map[string]SubCommand) Command {
	return func(args []string, procInout ProcInout) byte {
		flags := flag.NewFlagSet(name, flag.ContinueOnError)
		flags.SetOutput(procInout.Stderr)

		flags.Usage = func() {
			Usage(procInout.Stderr, name, subCommands)
		}

		if err := flags.Parse(args); err != nil {
			if errors.Is(err, flag.ErrHelp) {
				return 0
			}
			return 1
		}

		if len(flags.Args()) == 0 {
			Usage(procInout.Stderr, name, subCommands)
			return 1
		}

		subCommand, ok := subCommands[args[0]]
		if !ok {
			_, _ = procInout.Stderr.Write([]byte("unknown subcommand\n"))
			Usage(procInout.Stderr, name, subCommands)
			return 1
		}

		return subCommand.Command(args[1:], procInout)
	}
}
