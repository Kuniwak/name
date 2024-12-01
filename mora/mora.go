package mora

// SuteGanaExceptOneMora is a set of Sute-Gana ("捨て仮名") that not count as one mora.
var SuteGanaExceptOneMora = map[rune]struct{}{
	'ァ': {},
	'ィ': {},
	'ゥ': {},
	'ェ': {},
	'ォ': {},
	'ャ': {},
	'ュ': {},
	'ョ': {},
}

func Count(rs []rune) byte {
	var count byte
	for _, r := range rs {
		if _, ok := SuteGanaExceptOneMora[r]; ok {
			continue
		}
		count++
	}
	return count
}
