package pair

type Pair[F, S any] struct {
	First  F
	Second S
}

func New[F, S any](first F, second S) Pair[F, S] {
	return Pair[F, S]{First: first, Second: second}
}

func (p Pair[F, S]) Swap() Pair[S, F] {
	return Pair[S, F]{First: p.Second, Second: p.First}
}

func Map[F1, F2, S1, S2 any](p1 Pair[F1, S1], f func(F1, S1) (F2, S2)) Pair[F2, S2] {
	return New(f(p1.First, p1.Second))
}

func ToMap[K comparable, V any](ps []Pair[K, V]) map[K]V {
	m := make(map[K]V, len(ps))
	for _, p := range ps {
		m[p.First] = p.Second
	}
	return m
}

func FromMap[K comparable, V any](m map[K]V) []Pair[K, V] {
	pairs := make([]Pair[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, New(k, v))
	}
	return pairs
}

func FirstSlice[F, S any](ps []Pair[F, S]) []F {
	fs := make([]F, len(ps))
	for i, p := range ps {
		fs[i] = p.First
	}
	return fs
}

func SecondSlice[F, S any](ps []Pair[F, S]) []S {
	ss := make([]S, len(ps))
	for i, p := range ps {
		ss[i] = p.Second
	}
	return ss
}
