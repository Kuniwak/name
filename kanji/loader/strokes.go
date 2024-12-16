package loader

import (
	"encoding/json"
	"fmt"
	"github.com/Kuniwak/name/kanji/types"
	"golang.org/x/text/unicode/norm"
)

func LoadStrokes(bs []byte) map[rune]byte {
	var strokes []types.StrokeEntry
	if err := json.Unmarshal(bs, &strokes); err != nil {
		panic(fmt.Sprintf("failed to unmarshal strokes: %v", err))
	}

	m := make(map[rune]byte)
	for _, s := range strokes {
		r := []rune(norm.NFC.String(s.Kanji))[0]
		m[r] = s.Strokes
	}
	return m
}
