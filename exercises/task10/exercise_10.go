package main

import (
	"fmt"
	"sync"
)

// 10. Worker Pool: ограниченная параллельность
// Это классическая конструкция, которую часто просят реализовать на собеседованиях.
// Ничего сложного, но важно понимать и запомнить, как устроена синхронизация и как управлять количеством одновременно работающих горутин.
// Задача:
// Реализуйте воркер-пул, который обрабатывает список простых чисел с ограниченным количеством параллельных горутин.
// В main у вас есть список простых чисел от 1 до 20.
// Запустите 3 воркера (это горутины), каждый из которых будет обрабатывать одно число из списка.
// Каждый воркер должен просто напечатать число, которое он обрабатывает.
// Как только все числа будут обработаны, программа должна завершиться.
// Пример вывода программы:
// Worker 1 is processing number 1
// Worker 2 is processing number 2
// Worker 3 is processing number 3
// ...

//func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
//	...
//}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is processing number %d\n", id, job)
	}
}

func main() {
	const (
		numJobs    = 20
		numWorkers = 3
	)

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}
