package set

type Set[T comparable] map[T]struct{}

func New[T comparable](elems ...T) Set[T] {
	s := make(Set[T], len(elems))
	for _, e := range elems {
		s[e] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(elems ...T) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

func (s Set[T]) Remove(elems ...T) {
	for _, e := range elems {
		delete(s, e)
	}
}

func (s Set[T]) Filter(f func(T) bool) Set[T] {
	s0 := New[T]()
	for e := range s {
		if f(e) {
			s0.Add(e)
		}
	}
	return s0
}

func (s Set[T]) ContainsAny(es ...T) bool {
	for _, e := range es {
		if _, ok := s[e]; ok {
			return true
		}
	}
	return false
}

func (s Set[T]) ContainsAll(es ...T) bool {
	for _, e := range es {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) ToSlice() []T {
	elems := make([]T, 0, len(s))
	for e := range s {
		elems = append(elems, e)
	}
	return elems
}

func (s Set[T]) Equal(s0 Set[T]) bool {
	if len(s) != len(s0) {
		return false
	}
	for e := range s {
		if _, ok := s0[e]; !ok {
			return false
		}
	}
	return true
}
