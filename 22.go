package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	// Создание длинных чисел
	number1 := new(big.Int)
	number2 := new(big.Int)

	// Инициализация длинных чисел
	number1.MulRange(1, 40)
	number2.MulRange(1, 38)

	fmt.Printf("число 1 = %s\n", number1.Text(10))
	fmt.Printf("число 2 = %s\n", number2.Text(10))

	fmt.Printf("сумма: %s\n", Add(number1, number2))
	fmt.Printf("разность: %s\n", Sub(number1, number2))
	fmt.Printf("умножение: %s\n", Mul(number1, number2))
	fmt.Printf("деление: %s", Div(number1, number2))
}

func Add(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Add(num1, num2).Text(10)
}

func Sub(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Sub(num1, num2).Text(10)
}

func Mul(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Mul(num1, num2).Text(10)
}

func Div(num1 *big.Int, num2 *big.Int) string {
	res := new(big.Int)
	return res.Div(num1, num2).Text(10)
}
