package sliceutil

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	ss := [][]int{{1, 2}, {3, 4, 5}, {6}}
	actual := Flatten(ss)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(actual, expected) {
		t.Error(cmp.Diff(expected, actual))
	}
}
