package wrap

import (
	"sync"
)

// Mutex contains the value, each operation on value should be done in Mutex
// parameter T should be reference type(map, pointer),
// otherwise the mutation will not affect on the value
type Mutex[T any] interface {
	Mutation(func(T)) // Mutation use Write lock
	Query(func(T))    // Query use Read lock
	Unwrap() T        // Unwrap the value will return value using ReadLock
	Wrap(T)           // Wrap the value will replace the value
}

type RwLocker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

func NewMutex[T any](value T, lockers ...RwLocker) Mutex[T] {
	w := &wrap[T]{
		value: value,
	}
	if len(lockers) == 0 || lockers[0] == nil {
		w.lock = &sync.RWMutex{}
	} else {
		w.lock = lockers[0]
	}
	return w
}

type wrap[T any] struct {
	value T
	lock  RwLocker
}

func (w *wrap[T]) Mutation(fn func(T)) {
	w.lock.Lock()
	defer w.lock.Unlock()
	fn(w.value)
}

func (w *wrap[T]) Query(fn func(T)) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	fn(w.value)
}

func (w *wrap[T]) Unwrap() T {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.value
}

func (w *wrap[T]) Wrap(t T) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.value = t
}

func NewRwLocker(locker sync.Locker) RwLocker {
	if locker == nil {
		panic("[NewRwLocker] locker must not be nil")
	}
	return &rwLock{
		locker: locker,
	}
}

type rwLock struct {
	locker sync.Locker
}

func (r *rwLock) Lock() {
	r.locker.Lock()
}

func (r *rwLock) Unlock() {
	r.locker.Unlock()
}

func (r *rwLock) RLock() {
	r.Lock()
}

func (r *rwLock) RUnlock() {
	r.Unlock()
}

var (
	_ Mutex[struct{}] = &wrap[struct{}]{}
	_ RwLocker        = &rwLock{}
)
