package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func checkKeyArray(arr []int, elem int) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}
	return false
}

func cross(set1 []int, set2 []int) (result []int) {
	if len(set2) < len(set1) {
		set1, set2 = set2, set1
	}
	for _, elem := range set1 {
		if checkKeyArray(set2, elem) && !checkKeyArray(result, elem) {
			result = append(result, elem)
		}
	}
	return
}

func main() {
	numbers1 := []int{6, 6, 6, 13, 666, 4, 5, 3, 2}
	numbers2 := []int{1, 2, 8, 7, 12, 5, 1}

	fmt.Println("numbers 1 :", numbers1)
	fmt.Println("numbers 2 :", numbers2)

	fmt.Println("Пересечение :", cross(numbers1, numbers2))
}
