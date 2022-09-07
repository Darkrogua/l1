package main

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел
	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numbersSquare := []int{}
	wg := new(sync.WaitGroup)
	c := make(chan int, len(numbers))
	defer close(c)

	for _, n := range numbers {
		wg.Add(1)

		go getSquareNumber(n, wg, c)
	}

	wg.Wait()

	for range numbers {
		numbersSquare = append(numbersSquare, <-c)
	}

	fmt.Println(numbersSquare)
	for _, n := range numbersSquare {
		fmt.Println(n, ",")
	}

}

func getSquareNumber(n int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	c <- n * n

	//fmt.Printf("Горутина с числом %d начала выполнение \n", n)

	//numbersSquare = append(c, n*n) - тупой вывод без форматирования

}
