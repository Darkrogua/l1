package main

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var sumSquareNumbers int

	wg := new(sync.WaitGroup)
	c := make(chan int, len(numbers))
	defer close(c)

	for _, n := range numbers {
		wg.Add(1)
		go getSquareNumber(n, wg, c)
	}

	wg.Wait()

	for range numbers {
		sumSquareNumbers += <-c
	}

	fmt.Printf("Результат: %d", sumSquareNumbers)
}

func getSquareNumber(n int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	c <- n * n

	//fmt.Printf("Горутина с числом %d начала выполнение \n", n)
	//numbersSquare = append(c, n*n) - тупой вывод без форматирования

}
