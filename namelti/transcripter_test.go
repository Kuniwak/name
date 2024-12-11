package namelti

import (
	"github.com/shogo82148/go-mecab"
	"reflect"
	"testing"
)

func TestTranscripter_Transcript(t *testing.T) {
	m, err := mecab.New(map[string]string{})
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
