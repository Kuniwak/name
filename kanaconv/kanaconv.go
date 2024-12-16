package kanaconv

import "unicode"

const (
	// http://www.unicodemap.org/range/62/Hiragana/
	hiraganaLo = 0x3041 // ぁ

	// http://www.unicodemap.org/range/63/Katakana/
	katakanaLo = 0x30a1 // ァ

	codeDiff = katakanaLo - hiraganaLo
)

func Htok(src []rune) []rune {
	dst := make([]rune, len(src))
	for i, r := range src {
		switch {
		case unicode.In(r, unicode.Hiragana):
			dst[i] = r + codeDiff
		default:
			dst[i] = r
		}
	}
	return dst
}
