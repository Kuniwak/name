package filter

func Mora(byteFunc ByteFunc) Func {
	return func(target Target) bool {
		return byteFunc(target.Mora)
	}
}
