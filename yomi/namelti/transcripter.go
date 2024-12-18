package namelti

import (
	"github.com/shogo82148/go-mecab"
	"strings"
)

type Transcripter struct {
	mecab   mecab.MeCab
	lattice mecab.Lattice
}

func NewTranscripter(m mecab.MeCab) (*Transcripter, error) {
	lattice, err := mecab.NewLattice()
	if err != nil {
		return nil, err
	}
	return &Transcripter{mecab: m, lattice: lattice}, nil
}

func (t *Transcripter) Transcript(s string, nbest int) ([]string, error) {
	t.lattice.Clear()
	t.lattice.AddRequestType(mecab.RequestTypeNBest)
	t.lattice.SetSentence(s)
	if err := t.mecab.ParseLattice(t.lattice); err != nil {
		return nil, err
	}

	res := make([]string, 0, nbest)
	i := nbest
	w := &strings.Builder{}
	for t.lattice.Next() && i > 0 {
		start := t.lattice.BOSNode()
		for node := start.Next(); node.Stat() != mecab.EOSNode; node = node.Next() {
			yomi, ok := ParseFeature7(node.Feature())
			if ok {
				w.WriteString(yomi)
			}
		}
		if w.Len() > 0 {
			res = append(res, w.String())
		}
		w.Reset()
		i--
	}

	return res, nil
}

func (t *Transcripter) Close() {
	t.mecab.Destroy()
}

func ParseFeature7(feature string) (string, bool) {
	features := strings.Split(feature, ",")
	if len(features) == 9 {
		return features[7], true
	}
	return "", false
}
