package future

import (
	"fmt"
	"github.com/AY7295/go-option"
	"sync"
)

// Future is a type that represents a value that will be available in the onceFuture.
type Future[T any] interface {
	Await() option.Option[T]
}

type Execution[T any] func() option.Option[T]

// Async creates a new onceFuture with the given fn function.
func Async[T any](exec Execution[T]) Future[T] {
	initPool()
	of := &onceFuture[T]{
		exec: exec,
	}
	of.data = option.None[T](coreProcessorsPool.Submit(of.execute))
	return of
}

func ToExecution[T any](fn func() (T, error)) Execution[T] {
	return option.WrapFn(fn)
}

type onceFuture[T any] struct {
	once sync.Once
	exec Execution[T]
	data option.Option[T]
}

// Await will wait for the onceFuture to be processed and return the result.
func (f *onceFuture[T]) Await() option.Option[T] {
	f.execute()
	return f.data
}

// execute will run the fn function and store the result.
func (f *onceFuture[T]) execute() {
	f.once.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				f.data = option.None[T](fmt.Errorf("panic: %v", r))
			}
		}()

		f.data = f.exec()
	})
}
