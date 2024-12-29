package yomi

import (
	"flag"
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/mecabfactory"
	"github.com/Kuniwak/name/mecabfactory/dicdir"
	"github.com/Kuniwak/name/yomi"
	"io"
)

var SubCommand = cli.SubCommand{
	Help: "convert given names to yomi",
	Command: func(args []string, procInout cli.ProcInout) byte {
		flags := flag.NewFlagSet("yomi", flag.ContinueOnError)
		flags.SetOutput(procInout.Stderr)
		flags.Usage = func() {
			_, _ = procInout.Stderr.Write([]byte("Usage: name yomi [OPTIONS] <GIVEN_NAME>\n"))
			_, _ = io.WriteString(procInout.Stderr, `
OPTIONS`)
			flags.PrintDefaults()
			_, _ = io.WriteString(procInout.Stderr, `
EXAMPLES
	$ name yomi 太郎
	タロウ
	タロウ
	フトシロウ
	フトロウ
	タイロウ
`)
		}

		nBest := flags.Int("n", 5, "number of yomi (used as N-best for MeCab internally)")

		if err := flags.Parse(args); err != nil {
			return 1
		}

		if flags.NArg() == 0 {
			flags.Usage()
			return 1
		}

		givenName := []rune(flags.Arg(0))

		m, err := mecabfactory.WithDictionary(dicdir.FirstAvailable(dicdir.ByDictNamesWithSearchPaths(
			dicdir.SearchPathByOS(),
			dicdir.NeologdOrIPADicUTF8(),
		)))
		if err != nil {
			_, _ = procInout.Stderr.Write([]byte("failed to initialize MeCab: " + err.Error() + "\n"))
			return 1
		}

		getYomi, err := yomi.ByMeCab(m, *nBest)
		if err != nil {
			_, _ = procInout.Stderr.Write([]byte("failed to initialize Transcripter: " + err.Error() + "\n"))
			return 1
		}

		yomis, err := getYomi(givenName)
		if err != nil {
			_, _ = procInout.Stderr.Write([]byte("failed to get yomi: " + err.Error() + "\n"))
			return 1
		}

		newline := []byte("\n")
		for _, y := range yomis {
			_, _ = io.WriteString(procInout.Stdout, string(y.Runes))
			_, _ = procInout.Stdout.Write(newline)
		}

		return 0
	},
}
