package filter

type ByteFunc func(byte) bool

func ByteEqual(n byte) ByteFunc {
	return func(c byte) bool {
		return c == n
	}
}

func ByteLessThan(n byte) ByteFunc {
	return func(c byte) bool {
		return c < n
	}
}

func ByteGreaterThan(n byte) ByteFunc {
	return func(c byte) bool {
		return c > n
	}
}
