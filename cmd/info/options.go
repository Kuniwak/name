package info

import (
	"errors"
	"flag"
	"fmt"
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
}

func ParseOptions(args []string, stderr io.Writer, cm map[rune]struct{}) (Options, error) {
	flags := flag.NewFlagSet("info", flag.ContinueOnError)

	flags.SetOutput(stderr)

	flags.Usage = func() {
		o := flags.Output()
		_, _ = fmt.Fprintf(o, "Usage: name info [options] <familyName> <givenName> <yomi>\n")
		_, _ = fmt.Fprintf(o, `
EXAMPLES
	$ name info 山田 太郎 タロウ
	評点    画数    名前    読み    天格    地格    人格    外格    総格
	8       13      太郎    タロウ  吉      大吉    大凶    大凶    大吉
`)
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return Options{Help: true}, nil
		}
		return Options{}, err
	}

	if len(args) == 0 {
		return Options{}, errors.New("given name is required")
	}

	familyName := []rune(args[0])
	if len(familyName) == 0 {
		return Options{}, fmt.Errorf("family name is required")
	}

	if !kanji.IsValid(familyName, cm) {
		return Options{}, fmt.Errorf("invalid kanji included: %q", familyName)
	}

	givenName := []rune(args[1])
	if len(givenName) == 0 {
		return Options{}, fmt.Errorf("given name is required")
	}

	if !kanji.IsValid(givenName, cm) {
		return Options{}, fmt.Errorf("invalid kanji included: %q", givenName)
	}

	y := kanaconv.Htok([]rune(norm.NFC.String(args[2])))
	if len(y) == 0 {
		return Options{}, fmt.Errorf("yomi-gana is required")
	}

	return Options{
		FamilyName: familyName,
		GivenName:  givenName,
		Yomi:       y,
	}, nil
}
