package main

import (
	"fmt"
	"sync"
)

// 5. Пайплайн обработки данных
// Задача:
// Необходимо реализовать пайплайн из **трёх стадий** (три горутины):
// 1. Первая стадия: отправляет в канал числа от 1 до 5.
// 2. Вторая стадия: получает числа из первого канала и возводит их в квадрат.
// 3. Третья стадия: выводит полученные числа.
// Каждая стадия должна работать независимо, передавая данные между собой через каналы. Программа должна завершаться корректно, когда все данные обработаны.

func producer(ch1 chan int) {
	for i := range 5 {
		ch1 <- i + 1
	}
	close(ch1)
}

func square(ch1, ch2 chan int) {
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

func printer(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch2 {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(ch1)
	go square(ch1, ch2)
	go printer(ch2, &wg)

	wg.Wait()
	// ожидание завершения (sync.WaitGroup)
}
