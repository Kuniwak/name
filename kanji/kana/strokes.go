package kana

import (
	_ "embed"
	kana2 "github.com/Kuniwak/name/kanaconv"
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
		result[kana] = [][]rune{kana2.Htok([]rune{kana})}
	}
	return result
}
