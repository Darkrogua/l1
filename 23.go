package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	i := 5
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("срез до %d\n", test)
	test = append(test[0:i], test[i+1:]...)

	fmt.Printf("срез после %d\n", test)

}
