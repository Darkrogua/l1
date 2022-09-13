package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

*/

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	var mapString = make(map[string]int)

	for id, val := range strings {
		mapString[val] = id
	}
	fmt.Println(mapString)
}
