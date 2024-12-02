package kanji

import (
	"github.com/Kuniwak/name/kanji/jinmei"
	"github.com/Kuniwak/name/kanji/joyo"
)

func LoadStrokes() map[rune]byte {
	joyos := joyo.LoadStrokes()
	jins := jinmei.LoadStrokes()

	strokes := make(map[rune]byte)
	for k, v := range joyos {
		strokes[k] = v
	}
	for k, v := range jins {
		strokes[k] = v
	}

	return strokes
}

func LoadYomi() map[rune][][]rune {
	joyos := joyo.LoadYomi()
	jins := jinmei.LoadYomi()

	yomi := make(map[rune][][]rune)
	for k, v := range joyos {
		yomi[k] = v
	}
	for k, v := range jins {
		yomi[k] = v
	}

	return yomi
}
