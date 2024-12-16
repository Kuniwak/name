package filter

import (
	"encoding/json"
	"fmt"
	"github.com/Kuniwak/name/eval"
	"golang.org/x/text/unicode/norm"
)

type Data struct {
	And          *[]Data         `json:"and,omitempty"`
	Or           *[]Data         `json:"or,omitempty"`
	Not          *Data           `json:"not,omitempty"`
	MinRank      *eval.Rank      `json:"minRank,omitempty"`
	MinTotalRank *byte           `json:"minTotalRank,omitempty"`
	Mora         *ByteFuncData   `json:"mora,omitempty"`
	Strokes      *ByteFuncData   `json:"strokes,omitempty"`
	True         *Unit           `json:"true,omitempty"`
	False        *Unit           `json:"false,omitempty"`
	YomiCount    *YomiCountData  `json:"yomiCount,omitempty"`
	YomiMatch    *MatchFuncData  `json:"yomi,omitempty"`
	KanjiCount   *KanjiCountData `json:"kanjiCount,omitempty"`
	KanjiMatch   *MatchFuncData  `json:"kanji,omitempty"`
	CommonYomi   *Unit           `json:"commonYomi,omitempty"`
	Length       *ByteFuncData   `json:"length,omitempty"`
	Sex          *string         `json:"sex,omitempty"`
}

type Unit struct{}

type YomiCountData struct {
	Rune  string       `json:"rune"`
	Count ByteFuncData `json:"count"`
}

type YomiMatchData struct {
	Runes string        `json:"runes"`
	Match MatchFuncData `json:"match"`
}

type KanjiCountData struct {
	Rune  string       `json:"rune"`
	Count ByteFuncData `json:"count"`
}

type KanjiMatchData struct {
	Runes string        `json:"runes"`
	Match MatchFuncData `json:"match"`
}

type ByteFuncData struct {
	LessThan    *byte `json:"lessThan,omitempty"`
	Equal       *byte `json:"equal,omitempty"`
	GreaterThan *byte `json:"greaterThan,omitempty"`
}

type ByteSetFuncData struct {
	Contain        *byte `json:"contain,omitempty"`
	GreaterThanAll *byte `json:"greaterThanAll,omitempty"`
	LessThanAll    *byte `json:"lessThanAll,omitempty"`
}

type IntFuncData struct {
	LessThan    *int `json:"lessThan,omitempty"`
	Equal       *int `json:"equal,omitempty"`
	GreaterThan *int `json:"greaterThan,omitempty"`
}

type MatchFuncData struct {
	Equal     *string `json:"equal,omitempty"`
	StartWith *string `json:"startWith,omitempty"`
	EndWith   *string `json:"endWith,omitempty"`
	Contain   *string `json:"contain,omitempty"`
}

func Parse(bs []byte) (*Data, error) {
	var seed Data
	if err := json.Unmarshal(bs, &seed); err != nil {
		return nil, err
	}
	return &seed, nil
}

func Build(seed *Data) (Func, error) {
	if seed.And != nil {
		filters := make([]Func, len(*seed.And))
		for i, s := range *seed.And {
			f, err := Build(&s)
			if err != nil {
				return nil, err
			}
			filters[i] = f
		}
		return And(filters...), nil
	}

	if seed.Or != nil {
		filters := make([]Func, len(*seed.Or))
		for i, s := range *seed.Or {
			f, err := Build(&s)
			if err != nil {
				return nil, err
			}
			filters[i] = f
		}
		return Or(filters...), nil
	}

	if seed.Not != nil {
		f, err := Build(seed.Not)
		if err != nil {
			return nil, err
		}
		return Not(f), nil
	}

	if seed.MinRank != nil {
		return MinRank(*seed.MinRank), nil
	}

	if seed.Mora != nil {
		intFunc, err := BuildByteFunc(*seed.Mora)
		if err != nil {
			return nil, err
		}
		return Mora(intFunc), nil
	}

	if seed.Strokes != nil {
		countFunc, err := BuildByteFunc(*seed.Strokes)
		if err != nil {
			return nil, err
		}
		return Strokes(countFunc), nil
	}

	if seed.True != nil {
		return True(), nil
	}

	if seed.False != nil {
		return False(), nil
	}

	if seed.YomiCount != nil {
		countFunc, err := BuildByteFunc(seed.YomiCount.Count)
		if err != nil {
			return nil, err
		}
		return YomiCount([]rune(norm.NFC.String(seed.YomiCount.Rune))[0], countFunc), nil
	}

	if seed.YomiMatch != nil {
		matchFunc, err := BuildMatchFunc(*seed.YomiMatch)
		if err != nil {
			return nil, err
		}
		return YomiMatch(matchFunc), nil
	}

	if seed.CommonYomi != nil {
		return CommonYomi(), nil
	}

	if seed.KanjiCount != nil {
		countFunc, err := BuildByteFunc(seed.KanjiCount.Count)
		if err != nil {
			return nil, err
		}
		return KanjiCount([]rune(norm.NFC.String(seed.KanjiCount.Rune))[0], countFunc), nil
	}

	if seed.KanjiMatch != nil {
		matchFunc, err := BuildMatchFunc(*seed.KanjiMatch)
		if err != nil {
			return nil, err
		}
		return KanjiMatch(matchFunc), nil
	}

	if seed.MinTotalRank != nil {
		return MinTotalRank(*seed.MinTotalRank), nil
	}

	if seed.Length != nil {
		countFunc, err := BuildByteFunc(*seed.Length)
		if err != nil {
			return nil, err
		}
		return Length(countFunc), nil
	}

	if seed.Sex != nil {
		switch *seed.Sex {
		case "asexual":
			return Sex(Asexual), nil
		case "male":
			return Sex(Male), nil
		case "female":
			return Sex(Female), nil
		default:
			return nil, fmt.Errorf("unknown sex: %s", *seed.Sex)
		}
	}

	return nil, fmt.Errorf("empty data")
}

func BuildByteFunc(data ByteFuncData) (ByteFunc, error) {
	if data.LessThan != nil {
		return ByteLessThan(*data.LessThan), nil
	}

	if data.Equal != nil {
		return ByteEqual(*data.Equal), nil
	}

	if data.GreaterThan != nil {
		return ByteGreaterThan(*data.GreaterThan), nil
	}

	return nil, fmt.Errorf("empty count data")
}

func BuildMatchFunc(data MatchFuncData) (MatchFunc, error) {
	if data.Equal != nil {
		return MatchExactly([]rune(norm.NFC.String(*data.Equal))), nil
	}

	if data.StartWith != nil {
		return MatchStartsWith([]rune(norm.NFC.String(*data.StartWith))), nil
	}

	if data.EndWith != nil {
		return MatchEndsWith([]rune(norm.NFC.String(*data.EndWith))), nil
	}

	if data.Contain != nil {
		return MatchContains([]rune(norm.NFC.String(*data.Contain))), nil
	}

	return nil, fmt.Errorf("empty match data")
}
