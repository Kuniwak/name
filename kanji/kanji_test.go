package kanji

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	strokesMap := map[rune]byte{
		'a': 1,
		'b': 2,
	}

	yomiMap := map[rune][][]rune{
		'a': {},
		'c': {},
	}

	actual := Load(strokesMap, yomiMap)
	expected := map[rune]struct{}{
		'a': {},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Error(cmp.Diff(expected, actual))
	}
}

func TestLoadStrokes(t *testing.T) {
	if len(LoadStrokes()) == 0 {
		t.Errorf("want non-empty, got empty")
	}
}

func TestLoadYomi(t *testing.T) {
	if len(LoadYomi()) == 0 {
		t.Errorf("want non-empty, got empty")
	}
}
