package types

type StrokeEntry struct {
	Kanji   string `json:"kanji"`
	Strokes byte   `json:"strokes"`
}

type YomiEntry struct {
	Kanji string   `json:"kanji"`
	Yomi  []string `json:"yomi"`
}

type Kanji struct {
	Strokes byte
	Yomi    [][]rune
}
