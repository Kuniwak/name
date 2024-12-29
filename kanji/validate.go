package kanji

import (
	"fmt"
	"github.com/Kuniwak/name/errorutil"
)

func IsValid1(r rune, cm map[rune]struct{}) error {
	_, ok := cm[r]
	if !ok {
		return fmt.Errorf("'%c' is not in 常用漢字 or 人名用漢字 or ひらがな or カタカナ", r)
	}
	return nil
}

func IsValid(givenName []rune, cm map[rune]struct{}) error {
	errs := make([]error, 0, len(givenName))
	for _, r := range givenName {
		if err := IsValid1(r, cm); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errorutil.Errors(errs)
	}
	return nil
}
