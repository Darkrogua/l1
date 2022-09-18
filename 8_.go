package main

import (
	"fmt"
	"math"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// Получает значение заданного бита в числе
func getBit(num int64, bit int) int64 {
	if bit > 0 {
		return (num >> (bit - 1)) % 2
	} else {
		fmt.Errorf("Wrong bit position")
		return 0
	}
}

// Устанавливает значение заданного бита в числе
func bitSetter(num int64, bit int, val bool) int64 {
	mask := int64(math.Pow(2, float64(bit-1))) // Выделяем нужный бит
	fmt.Println("mask", mask)

	if getBit(num, bit) == 1 && !val {
		return num &^ mask // Сброс единицы в ноль
	} else if getBit(num, bit) == 0 && val {
		return mask | num // Установка нуля в единицу
	} else {
		return num // Только два случая выше меняют значение исходного числа
	}
}

func main() {
	var num int64
	fmt.Println("Введите натуральное число")
	for num < 1 {
		fmt.Scan(&num)
	}

	// То что должно получиться смотреть в комментариях и сравнить с выводом
	fmt.Println(bitSetter(num, 1, true))  // 5 = 101
	fmt.Println(bitSetter(num, 1, false)) // 4 = 100
	fmt.Println(bitSetter(num, 2, true))  // 7 = 111
	fmt.Println(bitSetter(num, 2, false)) // 5 = 101
	fmt.Println(bitSetter(num, 3, true))  // 5 = 101
	fmt.Println(bitSetter(num, 3, false)) // 1 = 001
}
