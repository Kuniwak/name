package gen

type GenerateFunc func(familyName []rune, ch chan<- Generated)

type Generated struct {
	GivenName  []rune
	Yomi       []rune
	YomiString string
}
