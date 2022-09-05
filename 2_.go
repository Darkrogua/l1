package main

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел
	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	for _, n := range numbers {
		fmt.Println(n)
	}

}
