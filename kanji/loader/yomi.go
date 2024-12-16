package loader

import (
	"encoding/json"
	"fmt"
	"github.com/Kuniwak/name/kanji/types"
	"golang.org/x/text/unicode/norm"
)

func LoadYomi(bs []byte) map[rune][][]rune {
	var yomis []types.YomiEntry
	if err := json.Unmarshal(bs, &yomis); err != nil {
		panic(fmt.Sprintf("failed to unmarshal yomi: %v", err))
	}

	m := make(map[rune][][]rune)
	for _, y := range yomis {
		r := []rune(norm.NFC.String(y.Kanji))[0]

		rs := make([][]rune, len(y.Yomi))
		for i, yomi := range y.Yomi {
			rs[i] = []rune(norm.NFC.String(yomi))
		}
		m[r] = rs
	}
	return m
}
