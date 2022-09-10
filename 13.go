package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	var a, b = 23, 15
	a, b = b, a
	fmt.Print(a, b)
}
