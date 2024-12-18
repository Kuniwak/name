package mecabfactory

import (
	"github.com/Kuniwak/name/mecabfactory/dicdir"
	"github.com/shogo82148/go-mecab"
)

func WithDictionary(getDicDir dicdir.Func) (mecab.MeCab, error) {
	dicDir, err := getDicDir()
	if err != nil {
		return mecab.MeCab{}, err
	}
	return mecab.New(map[string]string{"dicdir": dicDir})
}
