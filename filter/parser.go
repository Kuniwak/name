package filter

import (
	"encoding/json"
	"fmt"
	"github.com/Kuniwak/name/eval"
)

type Data struct {
	And          *[]Data        `json:"and,omitempty"`
	Or           *[]Data        `json:"or,omitempty"`
	Not          *Data          `json:"not,omitempty"`
	MinRank      *eval.Rank     `json:"minRank,omitempty"`
	MinTotalRank *byte          `json:"minTotalRank,omitempty"`
	Mora         *ByteFuncData  `json:"mora,omitempty"`
	Strokes      *ByteFuncData  `json:"strokes,omitempty"`
	True         *Unit          `json:"true,omitempty"`
	False        *Unit          `json:"false,omitempty"`
	YomiCount    *YomiCountData `json:"yomiCount,omitempty"`
	CommonYomi   *Unit          `json:"commonYomi,omitempty"`
}

type Unit struct{}

type YomiCountData struct {
	Rune  string       `json:"rune"`
	Count ByteFuncData `json:"count"`
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

func Parse(bs []byte, yomiMap map[rune][][]rune) (Func, error) {
	var seed Data
	if err := json.Unmarshal(bs, &seed); err != nil {
		return nil, err
	}
	return Build(seed, yomiMap)
}

func Build(seed Data, yomiMap map[rune][][]rune) (Func, error) {
	if seed.And != nil {
		filters := make([]Func, len(*seed.And))
		for i, s := range *seed.And {
			f, err := Build(s, yomiMap)
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
			f, err := Build(s, yomiMap)
			if err != nil {
				return nil, err
			}
			filters[i] = f
		}
		return Or(filters...), nil
	}

	if seed.Not != nil {
		f, err := Build(*seed.Not, yomiMap)
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
		return YomiCount([]rune(seed.YomiCount.Rune)[0], countFunc), nil
	}

	if seed.CommonYomi != nil {
		return CommonYomi(), nil
	}

	if seed.MinTotalRank != nil {
		return MinTotalRank(*seed.MinTotalRank), nil
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
