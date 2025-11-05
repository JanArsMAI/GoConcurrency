package main

import (
	"fmt"
	"sync"
)

// 4. Передача данных по каналу
// Задача:
// Представьте, что у вас есть передатчик (генератор чисел) и приёмник (читатель).
// Передатчик генерирует числа от 1 до 5 и отправляет их по каналу.
// Приёмник читает эти числа и выводит их в консоль.
// После того как все числа отправлены, необходимо корректно закрыть канал.
// Реализуйте программу с двумя горутинами:
//   - одна пишет числа в канал
//   - вторая читает их и печатает

func generator(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		ch <- i + 1
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go generator(ch, &wg)
	wg.Add(1)
	go consumer(ch, &wg)
	wg.Wait()
}
