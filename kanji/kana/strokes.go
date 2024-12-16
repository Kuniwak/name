package kana

import (
	_ "embed"
	"github.com/Kuniwak/name/kanji/loader"
)

//go:embed data/strokes.json
var strokesBytes []byte

func LoadStrokes() map[rune]byte {
	return loader.LoadStrokes(strokesBytes)
}

func LoadYomi() map[rune][][]rune {
	strokes := LoadStrokes()
	result := make(map[rune][][]rune)
	for kana := range strokes {
		result[kana] = [][]rune{Htok([]rune{kana})}
	}
	return result
}