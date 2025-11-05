package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 2. Параллельный счётчик: Усложнение (atomic)
// Задача:
// В прошлой задаче мы использовали `sync.Mutex` для защиты общей переменной от гонки данных.
// Теперь реши ту же задачу с помощью **пакета `sync/atomic`, без использования мьютекса.
func main() {
	var counter atomic.Int64
	counter.Store(0)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("Result:", counter.Load())
}
