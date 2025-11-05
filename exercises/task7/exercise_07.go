package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 7. Гонки горутин
// Задача:
// Запустите две горутины, каждая из которых имитирует разное время выполнения (например, с помощью `time.Sleep`).
// Необходимо получить результат от той, которая завершится первой, вывести его и отменить оставшуюся горутину.

func work(ctx context.Context, name string, result chan string, delay time.Duration) {
	select {
	case <-time.After(delay):
		select {
		case result <- fmt.Sprintf("%s finished after %v", name, delay):
		case <-ctx.Done():
			return
		}
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan string, 1)

	go work(ctx, "Worker 1", resultCh, time.Duration(rand.Intn(500)+300)*time.Millisecond)
	go work(ctx, "Worker 2", resultCh, time.Duration(rand.Intn(500)+300)*time.Millisecond)

	result := <-resultCh
	fmt.Println("Result:", result)

	cancel()

	time.Sleep(100 * time.Millisecond)
}
