package maps

func Map[K1, K2 comparable, V1, V2 any](m1 map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	m2 := make(map[K2]V2, len(m1))
	for k1, v1 := range m1 {
		k2, v2 := f(k1, v1)
		m2[k2] = v2
	}
	return m2
}

func ForEach[K comparable, V any](m map[K]V, f func(K, V)) {
	for k, v := range m {
		f(k, v)
	}
}

func Filter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	m0 := make(map[K]V, len(m))
	for k, v := range m {
		if f(k, v) {
			m0[k] = v
		}
	}
	return m0
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
