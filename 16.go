package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{1, 8, 8, 9, 9, 13, -32, -2, 2}
	fmt.Println("Стартовый массив:\n", arr)
	quicksort(arr, 0, len(arr)-1)

	fmt.Println("Результат сортировки:\n", arr)
}

func quicksort(arr []int, first int, last int) {
	i := first
	j := last
	// Получаем элемент по середине
	p := arr[(i+j)/2]

	for {
		// Оставляем все что меньше слева
		for arr[i] < p {
			i++
		}

		// Оставляем все что меньше справа
		for arr[j] > p {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}

		if i >= j {
			break
		}
	}

	// Сортировка левой части
	if first < j {
		quicksort(arr, first, j)
	}

	// Сортировка правой части
	if i < last {
		quicksort(arr, i, last)
	}
}
