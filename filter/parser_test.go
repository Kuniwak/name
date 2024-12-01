package filter

import (
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/gen"
	"github.com/Kuniwak/name/kanji"
	"testing"
)

func TestParse(t *testing.T) {
	yomiMap := kanji.LoadYomi()
	strokesMap := kanji.LoadStrokes()

	tests := map[string]struct {
		input    string
		expected Func
	}{
		"True": {
			input:    `{"true":{}}`,
			expected: True(),
		},
		"False": {
			input:    `{"false":{}}`,
			expected: False(),
		},
		"And": {
			input:    `{"and":[]}`,
			expected: And(),
		},
		"Or": {
			input:    `{"or":[]}`,
			expected: Or(),
		},
		"Not": {
			input:    `{"not":{"false":{}}}`,
			expected: Not(False()),
		},
		"MinRank": {
			input:    `{"minRank":1}`,
			expected: MinRank(eval.Kyo),
		},
		"MinTotalRank": {
			input:    `{"minTotalRank":1}`,
			expected: MinTotalRank(1),
		},
		"MoraEqual": {
			input:    `{"mora":{"equal":3}}`,
			expected: Mora(ByteEqual(3)),
		},
		"MoraLessThan": {
			input:    `{"mora":{"lessThan":1}}`,
			expected: Mora(ByteLessThan(1)),
		},
		"MoraGreaterThan": {
			input:    `{"mora":{"greaterThan":1}}`,
			expected: Mora(ByteGreaterThan(1)),
		},
		"MaxStrokes": {
			input:    `{"strokes":{"greaterThan":1}}`,
			expected: Strokes(ByteGreaterThan(1)),
		},
		"YomiCountEqual": {
			input:    `{"yomiCount":{"rune":"タ","count":{"equal":1}}}`,
			expected: YomiCount('タ', ByteEqual(1)),
		},
		"YomiCountGreaterThan": {
			input:    `{"yomiCount":{"rune":"タ","count":{"greaterThan":1}}}`,
			expected: YomiCount('タ', ByteGreaterThan(1)),
		},
		"YomiCountLessThan": {
			input:    `{"yomiCount":{"rune":"タ","count":{"lessThan":1}}}`,
			expected: YomiCount('タ', ByteLessThan(1)),
		},
		"CommonYomi": {
			input:    `{"commonYomi":{}}`,
			expected: CommonYomi(),
		},
	}

	familyNames := []string{"山田", "一"}
	generateds := []gen.Generated{
		{
			GivenName:  []rune("太郎"),
			Yomi:       []rune("タロウ"),
			YomiString: "タロウ",
		},
		{
			GivenName:  []rune("花子"),
			Yomi:       []rune("ハナコ"),
			YomiString: "ハナコ",
		},
		{
			GivenName:  []rune("一"),
			Yomi:       []rune("ハジメ"),
			YomiString: "ハジメ",
		},
		{
			GivenName:  []rune("太郎太郎"),
			Yomi:       []rune("タロウタロウ"),
			YomiString: "タロウタロウ",
		},
	}

	for name, c := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Parse([]byte(c.input), yomiMap)
			if err != nil {
				t.Error(err.Error())
				return
			}

			for _, familyName := range familyNames {
				for _, generated := range generateds {
					evalResult, err := eval.Evaluate([]rune(familyName), generated.GivenName, strokesMap)
					if err != nil {
						t.Error(err.Error())
						return
					}

					target := Target{
						GivenName:  generated.GivenName,
						Yomi:       generated.Yomi,
						YomiString: generated.YomiString,
						Strokes:    eval.SumStrokes(generated.GivenName, strokesMap),
						EvalResult: evalResult,
					}
					if actual(target) != c.expected(target) {
						t.Errorf("expected: %v, actual: %v at %v", c.expected(target), actual(target), target)
					}
				}
			}
		})
	}
}
