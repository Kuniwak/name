package search

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/mecabfactory"
	"github.com/Kuniwak/name/mecabfactory/dicdir"
	"github.com/Kuniwak/name/strokes"
	"github.com/Kuniwak/name/yomi"
	"io"
	"os"
)

type Options struct {
	Help          bool
	Filter        filter.Func
	FamilyName    []rune
	MinLength     int
	MaxLength     int
	GeneratorFunc gen.GenerateFunc
}

func ParseOptions(args []string, stdin io.Reader, stderr io.Writer, cm map[rune]struct{}, strokesFunc strokes.Func, yomiFunc yomi.Func) (*Options, error) {
	flags := flag.NewFlagSet("search", flag.ContinueOnError)

	flags.SetOutput(stderr)
	space := flags.String("space", "common", "Search spaces (available: full, common)")
	minLength := flags.Int("min-length", 1, "Minimum length of a given name")
	maxLength := flags.Int("max-length", 3, "Maximum length of a given name")
	unsafeYomiCount := flags.Int("yomi-count", 5, "Number of Yomi-Gana candidates")
	unsafeDicDir := flags.String("dir-dict", "", "Directory of MeCab dictionary (full space only)")

	flags.Usage = func() {
		o := flags.Output()
		_, _ = fmt.Fprintf(o, "Usage: name [options] <familyName>\n\n")
		_, _ = fmt.Fprintf(o, "OPTIONS\n")
		flags.PrintDefaults()
		_, _ = fmt.Fprintf(o, `
STDIN
	Filter notated in JSON. See "name filter validate --help" for details.
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

	_, err := strokes.Sum(familyName, strokesFunc)
	if err != nil {
		return nil, fmt.Errorf("invalid family name: %w", err)
	}

	var genFunc gen.GenerateFunc
	switch *space {
	case "full":
		yomiCount := *unsafeYomiCount
		if yomiCount < 1 {
			return nil, errors.New("yomi-count must be greater than or equal to 1")
		}

		var dicDirFunc dicdir.Func
		if len(*unsafeDicDir) == 0 {
			dicDirFunc = dicdir.Fallback(dicdir.Neologd(dicdir.ByMecabConfig()), dicdir.Ipa(dicdir.ByMecabConfig()))
		} else {
			stat, err := os.Stat(*unsafeDicDir)
			if err != nil {
				return nil, fmt.Errorf("failed to stat %q: %w", *unsafeDicDir, err)
			}

			if !stat.IsDir() {
				return nil, fmt.Errorf("%q is not a directory", *unsafeDicDir)
			}
			dicDir := *unsafeDicDir
			dicDirFunc = dicdir.ByConstant(dicDir, nil)
		}

		m, err := mecabfactory.WithDictionary(dicDirFunc)
		if err != nil {
			return nil, fmt.Errorf("failed to create MeCab: %w", err)
		}

		yomiFuncByMeCab, err := yomi.ByMeCab(m, yomiCount)
		if err != nil {
			return nil, err
		}

		genFunc, err = gen.NewFullSpaceGenerator(cm, strokesFunc, yomi.Fallback(yomiFuncByMeCab, yomiFunc))
		if err != nil {
			return nil, err
		}
	case "common":
		genFunc = gen.NewCommonSpaceGenerator(cm)
	default:
		return nil, fmt.Errorf("unknown space: %q", *space)
	}

	bs, err := io.ReadAll(stdin)
	if err != nil {
		return nil, err
	}

	data, err := filter.Parse(bs)
	if err != nil {
		return nil, err
	}

	f, err := filter.Build(data)
	if err != nil {
		return nil, err
	}

	if *minLength < 1 {
		return nil, errors.New("min-length must be greater than or equal to 1")
	}

	if *minLength > 4 {
		return nil, errors.New("min-length must be less than 4")
	}

	if *maxLength > 4 {
		return nil, errors.New("max-length must be less than 4")
	}

	if *minLength > *maxLength {
		return nil, errors.New("min-length must be less than or equal to max-length")
	}

	return &Options{
		Filter:        f,
		FamilyName:    familyName,
		GeneratorFunc: genFunc,
		MinLength:     *minLength,
		MaxLength:     *maxLength,
	}, nil
}
