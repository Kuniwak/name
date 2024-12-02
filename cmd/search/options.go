package search

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"io"
)

type Options struct {
	Help          bool
	Filter        filter.Func
	FamilyName    []rune
	GeneratorFunc gen.GenerateFunc
}

func ParseOptions(args []string, stdin io.Reader, stderr io.Writer, strokesMap map[rune]byte, yomiMap map[rune][][]rune) (*Options, error) {
	flags := flag.NewFlagSet("search", flag.ContinueOnError)

	flags.SetOutput(stderr)
	space := flags.String("space", "common", "Search spaces (available: full, common)")

	flags.Usage = func() {
		o := flags.Output()
		_, _ = fmt.Fprintf(o, "Usage: name [options] <familyName>\n\n")
		_, _ = fmt.Fprintf(o, "OPTIONS\n")
		flags.PrintDefaults()
		_, _ = fmt.Fprintf(o, `
STDIN
	See $ name filter try -h
`)

		_, _ = fmt.Fprintf(o, `
EXAMPLES
	$ name search 山田 < ./filter.example.json
	評点    画数    名前    読み    天格    地格    人格    外格    総格
	15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
	15      13      一喜    イッキ  吉      大吉    大吉    大大吉  大吉
	...
`)
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return &Options{Help: true}, nil
		}
		return nil, err
	}

	familyName := []rune(flags.Arg(0))
	if len(familyName) == 0 {
		return nil, errors.New("family name is required")
	}

	for _, r := range familyName {
		if _, ok := yomiMap[r]; !ok {
			return nil, fmt.Errorf("unknown rune: %c", r)
		}
	}

	var genFunc gen.GenerateFunc
	switch *space {
	case "full":
		genFunc = gen.NewFullSpaceGenerator(strokesMap, yomiMap)
	case "common":
		genFunc = gen.NewCommonSpaceGenerator(strokesMap)
	default:
		return nil, fmt.Errorf("unknown space: %q", *space)
	}

	bs, err := io.ReadAll(stdin)
	if err != nil {
		return nil, err
	}

	f, err := filter.Parse(bs)

	return &Options{
		Filter:        f,
		FamilyName:    familyName,
		GeneratorFunc: genFunc,
	}, nil
}
