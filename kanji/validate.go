package kanji

func IsValid(givenName []rune, cm map[rune]struct{}) bool {
	for _, r := range givenName {
		if _, ok := cm[r]; !ok {
			return false
		}
	}
	return true
}
