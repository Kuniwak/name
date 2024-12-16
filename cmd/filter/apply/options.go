package apply

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Kuniwak/name/filter"
	"io"
	"os"
)

type Options struct {
	Help       bool
	FamilyName []rune
	Filter     filter.Func
	Result     io.Reader
}

func ParseOptions(args []string, stdin io.Reader, stderr io.Writer) (Options, error) {
	flags := flag.NewFlagSet("apply", flag.ContinueOnError)
	flags.SetOutput(stderr)

	flags.Usage = func() {
		_, _ = stderr.Write([]byte("Usage: name filter apply <familyName> --to <path>\n"))
		_, _ = fmt.Fprintf(stderr, "OPTIONS\n")
		flags.PrintDefaults()
		_, _ = fmt.Fprintf(stderr, `
STDIN
	See $ name filter test -h

EXAMPLES
	$ name filter apply 山田 --to /path/to/result.tsv < ./filter.example.json 
	評点    画数    名前    読み    天格    地格    人格    外格    総格
	15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
	15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
`)
	}

	unsafeResultPath := flags.String("to", "", "path to the result file of `name search`")

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return Options{Help: true}, nil
		}
		return Options{}, err
	}

	if *unsafeResultPath == "" {
		return Options{}, errors.New("missing required option: --to")
	}
	result, err := os.OpenFile(*unsafeResultPath, os.O_RDONLY, 0644)
	if err != nil {
		return Options{}, fmt.Errorf("failed to open result file: %w", err)
	}

	filterString, err := io.ReadAll(stdin)
	if err != nil {
		return Options{}, fmt.Errorf("failed to read filter from stdin: %w", err)
	}
	filterData, err := filter.Parse(filterString)
	if err != nil {
		return Options{}, fmt.Errorf("failed to parse filter: %w", err)
	}
	filterFunc, err := filter.Build(filterData)
	if err != nil {
		return Options{}, fmt.Errorf("failed to build filter: %w", err)
	}

	return Options{
		FamilyName: []rune(flags.Arg(0)),
		Filter:     filterFunc,
		Result:     result,
	}, nil
}
