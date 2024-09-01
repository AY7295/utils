package future

import (
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

var (
	coreProcessorsPool *ants.Pool
	onceInitPool       sync.Once
	maxExecutingFuture = 10
)

func initPool() {
	onceInitPool.Do(func() {
		coreProcessorsPool = newAntsPool()
	})
}

func newAntsPool() *ants.Pool {
	pool, err := ants.NewPool(maxExecutingFuture,
		ants.WithPreAlloc(true),
		ants.WithMaxBlockingTasks(maxExecutingFuture*100),
		ants.WithExpiryDuration(time.Minute*10),
	)
	if err != nil {
		panic(err)
	}
	return pool
}

func SetMaxGoroutines(maxGoroutines int) {
	maxExecutingFuture = maxGoroutines
}
