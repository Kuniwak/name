package sliceutil

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestCartesian(t *testing.T) {
	tc := map[string]struct {
		input    [][]string
		expected [][]string
	}{
		"empty (outer)": {
			input:    [][]string{},
			expected: [][]string{},
		},
		"empty (inner)": {
			input: [][]string{
				{"a"},
				{},
				{"1"},
			},
			expected: [][]string{},
		},
		"single": {
			input: [][]string{
				{"a"},
				{"1"},
			},
			expected: [][]string{
				{"a", "1"},
			},
		},
		"several": {
			input: [][]string{
				{"a", "b"},
				{"1", "2", "3"},
			},
			expected: [][]string{
				{"a", "1"},
				{"a", "2"},
				{"a", "3"},
				{"b", "1"},
				{"b", "2"},
				{"b", "3"},
			},
		},
	}

	for name, c := range tc {
		t.Run(name, func(t *testing.T) {
			actual := Cartesian(c.input...)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Error(cmp.Diff(c.expected, actual))
			}
		})
	}
}
