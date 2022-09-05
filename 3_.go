package main

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	for _, n := range numbers {
		fmt.Println(n)
	}

}
