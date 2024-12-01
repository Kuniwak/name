package cmd

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

func ParseOptions(args []string, stdin io.Reader, stderr io.Writer, strokesMap map[rune]byte, yomiMap map[rune][][]rune) (*Options, func(), error) {
	flags := flag.NewFlagSet("name", flag.ContinueOnError)

	flags.SetOutput(stderr)
	space := flags.String("space", "", "Search spaces (available: full, common)")

	flags.Usage = func() {
		o := flags.Output()
		_, _ = fmt.Fprintf(o, "Usage: name [options] <familyName>\n\n")
		_, _ = fmt.Fprintf(o, "OPTIONS\n")
		flags.PrintDefaults()
		_, _ = fmt.Fprintf(o, `STDIN
	JSON filter

	true: {"true":{}}
	false: {"false":{}}
	and: {"and":[filter...]}
	or: {"or":[filter...]}
	not: {"not":filter}
	minRank: {"minRank":rank}
	rank: 0-4 (4=大大吉, 3=大吉, 2=吉, 1=凶, 0=大凶)
	minTotalRank: {"minTotalRank":byte}
	mora: {"maxMora":count}
	strokes: {"strokes":count}
	yomiCount: {"yomiCount":{"rune":string,"count":count}}
	count: {"equal":byte} or {"greaterThan":byte} or {"lessThan":byte}
`)
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return &Options{Help: true}, flags.Usage, nil
		}
		return nil, nil, err
	}

	familyName := []rune(flags.Arg(0))
	if len(familyName) == 0 {
		return nil, nil, errors.New("family name is required")
	}

	for _, r := range familyName {
		if _, ok := yomiMap[r]; !ok {
			return nil, nil, fmt.Errorf("unknown rune: %c", r)
		}
	}

	var genFunc gen.GenerateFunc
	switch *space {
	case "full":
		genFunc = gen.NewFullSpaceGenerator(strokesMap, yomiMap)
	case "common":
		genFunc = gen.NewCommonSpaceGenerator(strokesMap)
	default:
		return nil, nil, fmt.Errorf("unknown space: %q", *space)
	}

	bs, err := io.ReadAll(stdin)
	if err != nil {
		return nil, nil, err
	}

	f, err := filter.Parse(bs, yomiMap)

	return &Options{
		Filter:        f,
		FamilyName:    familyName,
		GeneratorFunc: genFunc,
	}, nil, nil
}
