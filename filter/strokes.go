package filter

func Strokes(countFunc ByteFunc) Func {
	return func(res Target) bool {
		return countFunc(res.Strokes)
	}
}
