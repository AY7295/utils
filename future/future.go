package future

import (
	"fmt"
	"github.com/AY7295/go-option"
	"sync"
)

// Future is a type that represents a value that will be available in the future.
type Future[T any] interface {
	Await() option.Option[T]
}

type Option[T any] func(f *future[T])

type Execution[T any] func() option.Option[T]

// Async creates a new future with the given fn function.
func Async[T any](exec Execution[T], opts ...Option[T]) Future[T] {
	return newWithOption(exec, opts...)
}

func ToExecution[T any](fn func() (T, error)) Execution[T] {
	return option.WrapFn(fn)
}

type future[T any] struct {
	sync.Once
	exec     Execution[T]
	data     option.Option[T]
	onDemand bool
}

// Await will wait for the future to be processed and return the result.
func (f *future[T]) Await() option.Option[T] {
	f.execute()
	return f.data
}

// newWithOption creates a new future with the given options.
func newWithOption[T any](exec Execution[T], opts ...Option[T]) Future[T] {
	f := &future[T]{
		exec: exec,
		data: option.None[T](),
	}
	for _, opt := range opts {
		opt(f)
	}

	if !f.onDemand {
		go f.execute()
	}
	return f
}

// execute will run the fn function and store the result.
func (f *future[T]) execute() {
	f.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				f.data = option.None[T](fmt.Errorf("panic: %v", r))
			}
		}()

		f.data = f.exec()
	})
}

// OnDemand is an option to create a future that will be processed when Await is called.
func OnDemand[T any]() Option[T] {
	return func(f *future[T]) {
		f.onDemand = true
	}
}
