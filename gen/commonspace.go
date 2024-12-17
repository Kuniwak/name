package gen

import (
	_ "embed"
	"encoding/json"
	"github.com/Kuniwak/name/kanji"
	"golang.org/x/text/unicode/norm"
)

//go:embed data/mei.json
var meiBytes []byte

type meiData map[string][]string

func NewCommonSpaceGenerator(cm map[rune]struct{}) GenerateFunc {
	var mei meiData
	if err := json.Unmarshal(meiBytes, &mei); err != nil {
		panic(err.Error())
	}

	return func(familyName []rune, opts Options, ch chan<- Generated) error {
		defer close(ch)
		for yomi, names := range mei {
			for _, name := range names {
				givenNameString := norm.NFC.String(name)
				givenName := []rune(givenNameString)

				if !kanji.IsValid(givenName, cm) || len(givenName) < opts.MinLength || len(givenName) > opts.MaxLength {
					continue
				}

				yomiString := norm.NFC.String(yomi)
				ch <- Generated{
					GivenName:  givenName,
					Yomi:       []rune(yomiString),
					YomiString: yomiString,
				}
			}
		}

		return nil
	}
}
