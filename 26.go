package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func main() {
	var x string
	var answer bool = true
	fmt.Scan(&x)
	lowStirng := strings.ToLower(x)
	rs := []rune(lowStirng)
	for i := range rs {
		if strings.Count(x, string(rs[i])) > 1 {
			answer = false
			break
		}
	}
	fmt.Print(answer)

}
