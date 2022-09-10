package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {

	var text string = "snow dog sun"
	fmt.Print(text, " - ")
	textWords := strings.Fields(text)
	for i := 0; i < len(textWords); i++ {
		fmt.Print(textWords[len(textWords)-(i+1)], " ")
	}

}

//func mirroкWord(s string) string {
//	runes := []rune(s)
//	runesOut := []rune(s)

//	for i := range runes {
//		runesOut[i] = runes[len(runes)-(i+1)]
//	}
//	return string(runesOut)
//}
