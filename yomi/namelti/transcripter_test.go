package namelti

import (
	"github.com/Kuniwak/name/mecabfactory"
	"github.com/Kuniwak/name/mecabfactory/dicdir"
	"reflect"
	"testing"
)

func TestTranscripter_Transcript(t *testing.T) {
	m, err := mecabfactory.WithDictionary(dicdir.FirstAvailable(dicdir.ByDictNamesWithSearchPaths(
		dicdir.SearchPathByOS(),
		dicdir.IPADicUTF8(),
	)))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	transcripter, err := NewTranscripter(m)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	defer transcripter.Close()

	ss, err := transcripter.Transcript("太郎", 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if !reflect.DeepEqual(ss, []string{"タロウ", "タロウ", "フトシロウ"}) {
		t.Errorf("want %v, but %v", []string{"タロウ", "タロ", "タロー"}, ss)
	}
}
