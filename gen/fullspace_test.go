package gen

import (
	"github.com/Kuniwak/name/kanji/loader"
	"github.com/Kuniwak/name/sliceutil"
	"github.com/Kuniwak/name/strokes"
	"github.com/Kuniwak/name/yomi"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/sync/errgroup"
	"reflect"
	"sort"
	"testing"
)

func TestNewFullSpaceGenerator(t *testing.T) {
	strokesMap := map[rune]byte{
		'太': 4,
		'郎': 9,
	}
	yomiMap := map[rune][][]rune{
		'太': {[]rune("タ"), []rune("タイ")},
		'郎': {[]rune("ロウ")},
	}
	cm := loader.Intersection(loader.Load(strokesMap), loader.Load(yomiMap))

	gen, err := NewFullSpaceGenerator(cm, strokes.ByMap(strokesMap), yomi.ByCartesian(yomiMap))
	if err != nil {
		t.Errorf("want nil, got %#v", err)
		return
	}

	var eg errgroup.Group
	ch := make(chan Generated)
	eg.Go(func() error { return gen([]rune("太郎"), Options{MinLength: 2, MaxLength: 2}, ch) })

	actual := sliceutil.FromChan(ch)
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].YomiString < actual[j].YomiString
	})

	expected := []Generated{
		{
			GivenName:  []rune("太太"),
			Yomi:       []rune("タイタ"),
			YomiString: "タイタ",
		},
		{
			GivenName:  []rune("太太"),
			Yomi:       []rune("タイタイ"),
			YomiString: "タイタイ",
		},
		{
			GivenName:  []rune("太郎"),
			Yomi:       []rune("タイロウ"),
			YomiString: "タイロウ",
		},
		{
			GivenName:  []rune("太太"),
			Yomi:       []rune("タタ"),
			YomiString: "タタ",
		},
		{
			GivenName:  []rune("太太"),
			Yomi:       []rune("タタイ"),
			YomiString: "タタイ",
		},
		{
			GivenName:  []rune("太郎"),
			Yomi:       []rune("タロウ"),
			YomiString: "タロウ",
		},
		{
			GivenName:  []rune("郎太"),
			Yomi:       []rune("ロウタ"),
			YomiString: "ロウタ",
		},
		{
			GivenName:  []rune("郎太"),
			Yomi:       []rune("ロウタイ"),
			YomiString: "ロウタイ",
		},
		{
			GivenName:  []rune("郎郎"),
			Yomi:       []rune("ロウロウ"),
			YomiString: "ロウロウ",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Error(cmp.Diff(expected, actual))
	}

	if err := eg.Wait(); err != nil {
		t.Errorf("want nil, got %#v", err)
	}
}
