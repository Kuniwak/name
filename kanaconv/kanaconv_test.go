package kanaconv

import "testing"

func TestHtok(t *testing.T) {
	testCases := map[string]struct {
		Input    string
		Expected string
	}{
		"empty": {
			Input:    "",
			Expected: "",
		},
		"ひらがな": {
			Input:    "ひらがな",
			Expected: "ヒラガナ",
		},
		"カタカナ": {
			Input:    "カタカナ",
			Expected: "カタカナ",
		},
		"漢字": {
			Input:    "漢字",
			Expected: "漢字",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := string(Htok([]rune(tc.Input)))
			if tc.Expected != actual {
				t.Errorf("want %q, got %q", tc.Expected, actual)
			}
		})
	}
}
