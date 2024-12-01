package mora

import "testing"

func TestCount(t *testing.T) {
	tcs := map[string]struct {
		input    string
		expected byte
	}{
		"Empty": {
			input:    "",
			expected: 0,
		},
		// SEE: https://ja.wikipedia.org/wiki/%E3%83%A2%E3%83%BC%E3%83%A9
		"猿": {
			input:    "サル",
			expected: 2,
		},
		"河童": {
			input:    "カッパ",
			expected: 3,
		},
		"チョコレート": {
			input:    "チョコレート",
			expected: 5,
		},
		"学校新聞": {
			input:    "ガッコウシンブン",
			expected: 8,
		},
		"学級新聞": {
			input:    "ガッキュウシンブン",
			expected: 8,
		},
		"観測": {
			input:    "カンソク",
			expected: 4,
		},
		"母さん": {
			input:    "カーサン",
			expected: 4,
		},
		"兄さん": {
			input:    "ニーサン",
			expected: 4,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			actual := Count([]rune(tc.input))
			if actual != tc.expected {
				t.Errorf("expected %d, but got %d at %q", tc.expected, actual, tc.input)
			}
		})
	}
}
