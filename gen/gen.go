package gen

type Options struct {
	MinLength int
	MaxLength int
}

type GenerateFunc func(familyName []rune, opts Options, ch chan<- Generated) error

type Generated struct {
	GivenName  []rune
	Yomi       []rune
	YomiString string
}
