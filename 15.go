package main

import "strings"

/*

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

/* у нас создается строка в 1024 символа, после мы из неё записываем в justString всего 100 символов ссылаясь на срез
тем самым мы используем только 100 символов, но в памяти будет хранится переменная v так как мы на неё ссылаемся и она не будет удалена
сборшиком мусора и у нас произойдет утечка памяти
*/

func main() {
	var justString string
	someFunc(&justString)
}

func someFunc(s *string) {
	//2 в 10й
	v := createHugeString(1 << 10)
	*s = strings.Clone(v)[:100]
}

func createHugeString(length int) string {
	r := ""
	for length > 0 {
		r += "qwerty"
		length -= 6
	}
	return r
}
