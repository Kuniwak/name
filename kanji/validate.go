package kanji

import (
	"github.com/Kuniwak/name/config"
)

func IsValid(givenName []rune, strokesMap map[rune]byte) bool {
	var acc byte = 0
	for _, r := range givenName {
		n, ok := strokesMap[r]
		if !ok {
			return false
		}
		acc += n
	}
	return acc <= config.MaxStrokes
}
