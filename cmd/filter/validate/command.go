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
STDIN
	Filter notated in JSON.

		filter       := true | false | and | or | not | sex | length | mora | strokes | minRank | minTotalRank |
                        yomiCount | yomi | kanjiCount | kanji
		true         := {"true": {}}
		false        := {"false": {}}
		and          := {"and": [filter...]}
		or           := {"or": [filter...]}
		not          := {"not": filter}
		sex          := {"sex": "asexual" | "male" | "female"}
		length       := {"length": count}
		mora         := {"maxMora": count}
		strokes      := {"strokes": count}
		minRank      := {"minRank": 0-4} (4=大大吉, 3=大吉, 2=吉, 1=凶, 0=大凶)
		minTotalRank := {"minTotalRank": byte}
		yomiCount    := {"yomiCount": {"rune": rune, "count": count}}
		yomi         := {"yomi": match}
		kanjiCount   := {"kanjiCount": {"rune": rune, "count": count}}
		kanji        := {"kanji": match}
		count        := {"equal": byte} | {"greaterThan": byte} | {"lessThan": byte}
		match        := {"equal": string} | {"contain": string} | {"startWith": string} | {"endWith": string}
		byte         := 0-255
		rune         := string that contains only one rune

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
