package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 3. Сумма чисел с распределением задач
// Задача:
// Разделите массив из 1000 случайных чисел на 10 равных частей.
// Запустите 10 горутин, каждая должна посчитать сумму своей части и отправить результат в канал.
// Главная функция должна собрать все результаты и вывести общую сумму.

func sumPart(nums []int, resultCh chan int) {
	ans := 0
	for _, v := range nums {
		ans += v
	}
	resultCh <- ans
}

func main() {
	const parts = 10
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	for i := 0; i < parts; i++ {
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			sumPart(nums[0+ind*100:100+ind*100], ch)
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	ans := 0
	for a := range ch {
		ans += a
	}
	fmt.Println(ans)
}
