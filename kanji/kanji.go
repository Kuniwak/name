package kanji

import (
	"github.com/Kuniwak/name/kanji/jinmei"
	"github.com/Kuniwak/name/kanji/joyo"
	"github.com/Kuniwak/name/kanji/kana"
	"github.com/Kuniwak/name/kanji/loader"
)

func Load(strokesMap, yomiMap map[rune]byte) map[rune]struct{} {
	return loader.Intersection2(loader.Load(strokesMap), loader.Load(yomiMap))
}

func LoadStrokes() map[rune]byte {
	joyos := joyo.LoadStrokes()
	jins := jinmei.LoadStrokes()
	kanas := kana.LoadStrokes()

	strokes := make(map[rune]byte)
	for k, v := range joyos {
		strokes[k] = v
	}
	for k, v := range jins {
		strokes[k] = v
	}
	for k, v := range kanas {
		strokes[k] = v
	}
	return strokes
}

func LoadYomi() map[rune][][]rune {
	joyos := joyo.LoadYomi()
	jins := jinmei.LoadYomi()
	kanas := kana.LoadYomi()

	yomi := make(map[rune][][]rune)
	for k, v := range joyos {
		yomi[k] = v
	}
	for k, v := range jins {
		yomi[k] = v
	}
	for k, v := range kanas {
		yomi[k] = v
	}
	return yomi
}
