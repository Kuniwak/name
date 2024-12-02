package filter

func Length(countFunc ByteFunc) Func {
	return func(res Target) bool {
		return countFunc(byte(len(res.GivenName)))
	}
}
