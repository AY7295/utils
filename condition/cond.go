package condition

func Choice[T any](cond bool, tv, fv T) T {
	if cond {
		return tv
	}
	return fv
}

// Lazy returns trueValue if condition is true, otherwise it returns the result of falseValue.
func Lazy[T any](cond bool, v T, onFalse func() T) T {
	if cond {
		return v
	}
	return onFalse()
}
