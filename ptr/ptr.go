package ptr

func Of[T any](t T) *T {
	return &t
}

func From[T any](p *T, onNil T) T {
	if p == nil {
		return onNil
	}
	return *p
}
