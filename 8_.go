package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var num int64
	fmt.Println("Введите натуральное число")
	for num < 1 {
		fmt.Scan(&num)
	}

}
