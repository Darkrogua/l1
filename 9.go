package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	s := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go sendNumber(numbers, c)
	go sQuadeNumber(c, s)
	go printData(s, &wg)

	wg.Wait()

}

func sQuadeNumber(c, q chan int) {
	for val := range c {
		q <- 2 * val
	}
	close(q)
}

func sendNumber(arr []int, c chan int) {
	for _, val := range arr {
		c <- val
	}
	close(c)
}

func printData(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range c {
		fmt.Println(val)
	}

}
