package kanji

import (
	"fmt"
	"github.com/Kuniwak/name/kanji/jinmei"
	"github.com/Kuniwak/name/kanji/joyo"
	"github.com/Kuniwak/name/kanji/types"
)

func LoadKanji() map[rune]*types.Kanji {
	strokes := LoadStrokes()
	yomis := LoadYomi()

	if len(strokes) != len(yomis) {
		panic(fmt.Sprintf("strokes and yomis have different length"))
	}

	kanji := make(map[rune]*types.Kanji)
	for k, v := range strokes {
		yomi, ok := yomis[k]
		if !ok {
			panic(fmt.Sprintf("yomi not found for %c", k))
		}

		kanji[k] = &types.Kanji{
			Strokes: v,
			Yomi:    yomi,
		}
	}

	return kanji
}

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
