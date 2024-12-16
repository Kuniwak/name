package filter

func KanjiCount(r rune, byteFunc ByteFunc) Func {
	return func(d Target) bool {
		var c byte = 0
		for _, r2 := range d.Kanji {
			if r == r2 {
				c++
			}
		}
		return byteFunc(c)
	}
}

func KanjiMatch(matchFunc MatchFunc) Func {
	return func(d Target) bool {
		return matchFunc(d.Kanji)
	}
}
