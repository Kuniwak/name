package filter

func And(filters ...Func) Func {
	return func(res Target) bool {
		for _, f := range filters {
			if !f(res) {
				return false
			}
		}
		return true
	}
}

func Or(filters ...Func) Func {
	return func(res Target) bool {
		for _, f := range filters {
			if f(res) {
				return true
			}
		}
		return false
	}
}

func Not(filter Func) Func {
	return func(res Target) bool {
		return !filter(res)
	}
}

func True() Func {
	return func(res Target) bool {
		return true
	}
}

func False() Func {
	return func(res Target) bool {
		return false
	}
}
