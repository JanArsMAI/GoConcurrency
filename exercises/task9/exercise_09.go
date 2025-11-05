package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 9. Атомарный счётчик
// Задача:
// Реализуйте счётчик, который увеличивается в нескольких горутинах без гонки данных. Используйте `sync/atomic`,
// чтобы обеспечить безопасный инкремент без мьютексов. После завершения всех горутин выведите итоговое значение счётчика.

func main() {
	var counter atomic.Int64
	counter.Store(0)
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(counter.Load())
}
