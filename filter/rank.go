package filter

import "github.com/Kuniwak/name/eval"

func MinRank(min eval.Rank) Func {
	return func(res Target) bool {
		// We cannot change Tenkaku so we don't need to check it.
		return res.EvalResult.Jinkaku >= min &&
			res.EvalResult.Chikaku >= min &&
			res.EvalResult.Gaikaku >= min &&
			res.EvalResult.Sokaku >= min
	}
}

func MinTotalRank(min byte) Func {
	return func(res Target) bool {
		return res.EvalResult.Total() >= min
	}
}
