package kanji

import (
	"fmt"
	"github.com/Kuniwak/name/kanji/jinmei"
	"github.com/Kuniwak/name/kanji/joyo"
	"github.com/Kuniwak/name/kanji/kana"
	"github.com/Kuniwak/name/kanji/loader"
)

func Load(strokesMap map[rune]byte, yomiMap map[rune][][]rune) map[rune]struct{} {
	return loader.Intersection(loader.Load(strokesMap), loader.Load(yomiMap))
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
		if _, ok := strokes[k]; ok {
			panic(fmt.Sprintf("duplicate key between joyo and jinmei: %q", string(k)))
		}
		strokes[k] = v
	}
	for k, v := range kanas {
		if _, ok := strokes[k]; ok {
			panic(fmt.Sprintf("duplicate key among joyo and jinmei and kana: %q", string(k)))
		}
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
		if _, ok := yomi[k]; ok {
			panic(fmt.Sprintf("duplicate key between joyo and jinmei: %q", string(k)))
		}
		yomi[k] = v
	}
	for k, v := range kanas {
		if _, ok := yomi[k]; ok {
			panic(fmt.Sprintf("duplicate key among joyo and jinmei and kana: %q", string(k)))
		}
		yomi[k] = v
	}
	return yomi
}
