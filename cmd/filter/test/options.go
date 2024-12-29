package test

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/kanaconv"
	"github.com/Kuniwak/name/kanji"
	"golang.org/x/text/unicode/norm"
	"io"
)

type Options struct {
	Help       bool
	FamilyName []rune
	GivenName  []rune
	Yomi       []rune
	Filter     filter.Func
}

func ParseOptions(args []string, stdin io.Reader, stderr io.Writer, cm map[rune]struct{}) (Options, error) {
	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	flags.SetOutput(stderr)

	flags.Usage = func() {
		_, _ = stderr.Write([]byte("Usage: name filter test <FAMILY_NAME> <GIVEN_NAME> <GIVEN_NAME_KANA>\n"))
		_, _ = fmt.Fprintf(stderr, `
STDIN
	Filter notated in JSON. See "name filter validate --help" for details.

EXAMPLES
	$ name filter test 田中 太郎 タロウ < filter.json
	$ echo $?
	0

	$ name filter test 田中 太郎 タロウ < filter.json
	$ echo $?
	1
`)
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return Options{Help: true}, nil
		}
		return Options{}, err
	}

	bs, err := io.ReadAll(stdin)
	if err != nil {
		return Options{}, err
	}

	data, err := filter.Parse(bs)
	if err != nil {
		return Options{}, err
	}

	f, err := filter.Build(data)
	if err != nil {
		return Options{}, err
	}

	familyName := []rune(flags.Arg(0))
	if len(familyName) == 0 {
		return Options{}, errors.New("family name is required")
	}

	if err := kanji.IsValid(familyName, cm); err != nil {
		return Options{}, err
	}

	givenName := []rune(flags.Arg(1))
	if len(givenName) == 0 {
		return Options{}, errors.New("given name is required")
	}

	if err := kanji.IsValid(givenName, cm); err != nil {
		return Options{}, err
	}

	y := kanaconv.Htok([]rune(norm.NFC.String(flags.Arg(2))))
	if len(y) == 0 {
		return Options{}, errors.New("yomigana is required")
	}

	return Options{
		FamilyName: familyName,
		GivenName:  givenName,
		Yomi:       y,
		Filter:     f,
	}, nil
}
