package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	arr := []int{2, 20, 40, 42, 49, 69, 96, 100}
	search := 96
	fmt.Println(arr)
	fmt.Printf("Поиск: %d, Индекс: %d", search, binarySearch(arr, search))
}

func binarySearch(arr []int, elem int) int {
	f := 0
	l := len(arr) - 1

	for {
		i := (f + l) / 2
		if arr[i] == elem {
			return i
		} else if elem < arr[i] {
			l = i - 1
		} else {
			f = i + 1
		}

		if f < 0 || l >= len(arr) {
			return -1
		}
	}
}
