package mecabfactory

import "github.com/shogo82148/go-mecab"

func WithDictionaryDirectory(dictDir string) (mecab.MeCab, error) {
	return mecab.New(map[string]string{"dicdir": dictDir})
}
