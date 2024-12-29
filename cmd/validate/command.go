package validate

import (
	"flag"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/kanji"
	"io"
)

var SubCommand = cli.SubCommand{
	Help: "validate given names",
	Command: func(args []string, procInout cli.ProcInout) byte {
		flags := flag.NewFlagSet("validate", flag.ContinueOnError)
		flags.SetOutput(procInout.Stderr)
		flags.Usage = func() {
			_, _ = procInout.Stderr.Write([]byte("Usage: name validate <GIVEN_NAME>\n"))
			_, _ = io.WriteString(procInout.Stderr, `
EXAMPLES
	$ name validate 太郎
	$ echo $?
	0

	$ name validate 龘
	'龘' is not in 常用漢字 or 人名用漢字 or ひらがな or カタカナ
	$ echo $?
	1
`)
		}

		if err := flags.Parse(args); err != nil {
			return 1
		}

		if flags.NArg() == 0 {
			flags.Usage()
			return 1
		}

		givenName := []rune(flags.Arg(0))

		if err := kanji.IsValid(givenName, kanji.Load(kanji.LoadStrokes(), kanji.LoadYomi())); err != nil {
			_, _ = io.WriteString(procInout.Stderr, err.Error())
			return 1
		}
		return 0
	},
}
