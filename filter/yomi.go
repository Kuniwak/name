package filter

import (
	_ "embed"
	"encoding/json"
	"golang.org/x/text/unicode/norm"
)

func YomiCount(r1 rune, byteFunc ByteFunc) Func {
	return func(d Target) bool {
		var c byte = 0
		for _, r2 := range d.Yomi {
			if r1 == r2 {
				c++
			}
		}
		return byteFunc(c)
	}
}

func YomiMatch(matchFunc MatchFunc) Func {
	return func(d Target) bool {
		return matchFunc(d.Yomi)
	}
}

//go:embed data/common.json
var commonYomisBytes []byte

func CommonYomi() Func {
	var commonYomis []string
	if err := json.Unmarshal(commonYomisBytes, &commonYomis); err != nil {
		panic(err)
	}

	sm := make(map[string]struct{})
	for _, yomi := range commonYomis {
		sm[norm.NFC.String(yomi)] = struct{}{}
	}

	return func(d Target) bool {
		_, ok := sm[d.YomiString]
		return ok
	}
}
