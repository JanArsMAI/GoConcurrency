package main

import (
	"context"
	"fmt"
	"time"
)

// 6. Остановка горутины через context.WithTimeout
// Задача:
// Запустите горутину, которая каждую секунду выводит `"tick"`.
// Используйте `context.WithTimeout`, чтобы горутина завершилась автоматически через 3 секунды.

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("tick")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
}
