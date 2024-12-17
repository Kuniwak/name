package yomi

import (
	"github.com/Kuniwak/name/kanji"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestByCartesian(t *testing.T) {
	yomiMap := kanji.LoadYomi()
	f := ByCartesian(yomiMap)

	actual, err := f([]rune("太郎"))
	if err != nil {
		t.Errorf("want nil, got %#v", err)
		return
	}

	expected := []Result{
		{[]rune("タイロウ"), "タイロウ"},
		{[]rune("タロウ"), "タロウ"},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Error(cmp.Diff(expected, actual))
	}
}
