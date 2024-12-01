package jinmei

import (
	_ "embed"
	"github.com/Kuniwak/name/kanji/loader"
)

//go:embed data/strokes.json
var strokesBytes []byte

//go:embed data/yomi.json
var yomiBytes []byte

func LoadStrokes() map[rune]byte {
	return loader.LoadStrokes(strokesBytes)
}

func LoadYomi() map[rune][][]rune {
	return loader.LoadYomi(yomiBytes)
}
