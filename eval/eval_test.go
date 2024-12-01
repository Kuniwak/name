package eval

import (
	"fmt"
	"github.com/Kuniwak/name/config"
	"github.com/Kuniwak/name/kanji"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := map[string]struct {
		familyName string
		givenName  string
		expected   Result
	}{
		"山田 太郎": {
			familyName: "山田",
			givenName:  "太郎",
			expected: Result{
				Tenkaku: Kichi,
				Jinkaku: DaiKyo,
				Chikaku: DaiKichi,
				Gaikaku: DaiKyo,
				Sokaku:  DaiKichi,
			},
		},
		"一 二": {
			familyName: "一",
			givenName:  "二",
			expected: Result{
				Tenkaku: DaiKichi,
				Jinkaku: DaiKichi,
				Chikaku: DaiKyo,
				Gaikaku: DaiKichi,
				Sokaku:  DaiKichi,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sm := kanji.LoadStrokes()

			actual, err := Evaluate([]rune(test.familyName), []rune(test.givenName), sm)
			if err != nil {
				t.Errorf("want nil, got %v", err)
				return
			}

			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf(cmp.Diff(test.expected, actual))
			}
		})
	}
}

func TestStrokesToRank(t *testing.T) {
	for i := byte(1); i <= config.MaxStrokes; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if _, err := StrokesToRank(i); err != nil {
				t.Errorf("want nil, got %v", err)
			}
		})
	}
}
