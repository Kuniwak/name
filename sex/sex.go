package sex

import (
	_ "embed"
	"encoding/json"
	"github.com/Kuniwak/name/kanaconv"
	"golang.org/x/text/unicode/norm"
)

//go:embed data/male.json
var maleBytes []byte

//go:embed data/female.json
var femaleBytes []byte

type Sex byte

const (
	Unknown Sex = 0
	Asexual Sex = 1
	Male    Sex = 2
	Female  Sex = 3

	UnknownString string = "不明"
	AsexualString string = "両性"
	MaleString    string = "男性"
	FemaleString  string = "女性"
)

func (s Sex) String() string {
	switch s {
	case Asexual:
		return AsexualString
	case Male:
		return MaleString
	case Female:
		return FemaleString
	}
	return UnknownString
}

type Func func(givenName string) Sex

func ByNameLists(maleNames, femaleNames map[string]struct{}) Func {
	return func(givenName string) Sex {
		_, mOk := maleNames[givenName]
		_, fOk := femaleNames[givenName]
		if mOk && fOk {
			return Asexual
		} else if mOk {
			return Male
		} else if fOk {
			return Female
		} else {
			return Unknown
		}
	}
}

func LoadMaleNames() map[string]struct{} {
	return LoadNames(maleBytes)
}

func LoadFemaleNames() map[string]struct{} {
	return LoadNames(femaleBytes)
}

func LoadNames(bs []byte) map[string]struct{} {
	var names []string
	if err := json.Unmarshal(bs, &names); err != nil {
		panic(err.Error())
	}

	normed := make(map[string]struct{}, len(names))
	for _, name := range names {
		yomi := kanaconv.Htok([]rune(norm.NFC.String(name)))
		normed[string(yomi)] = struct{}{}
	}

	return normed
}
