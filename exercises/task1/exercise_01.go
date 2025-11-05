package main

import (
	"fmt"
	"sync"
)

// 1. Параллельный счётчик
// Задача:
// Реализуйте программу, которая запускает 10 горутин, каждая из которых увеличивает общий счётчик на 1.
// Программа должна дождаться завершения всех горутин и вывести итоговое значение счётчика:
// Result: 10
// Важно: доступ к общей переменной должен быть потокобезопасным, чтобы избежать гонки данных.
func main() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("Result:", counter)
}
