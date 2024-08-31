package slice

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func ForEach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

func Filter[T any](s []T, f func(T) bool) []T {
	s0 := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			s0 = append(s0, v)
		}
	}
	return s0
}

func Reduce[T any, U any](s []T, f func(U, T) U, init U) U {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func ToMap[K comparable, T, V any](s []T, f func(T) (K, V)) map[K]V {
	m := make(map[K]V, len(s))
	for _, val := range s {
		k, v := f(val)
		m[k] = v
	}
	return m
}
