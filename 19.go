package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {

	var text string = "главрыба"

	runes := []rune(text)
	runesOut := []rune(text)

	for i := range runes {
		runesOut[i] = runes[len(runes)-(i+1)]
	}
	fmt.Println(string(runesOut))

}
