package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var tChanInt interface{} = make(chan int, 4)
	var tChanString interface{} = make(chan string)
	var tChanBool interface{} = make(chan bool)
	var tInt interface{} = 22
	var tString interface{} = "Hello Go"
	var tBool interface{} = true

	getTypeCustom(tChanInt)
	getTypeCustom(tChanString)
	getTypeCustom(tChanBool)
	getTypeCustom(tInt)
	getTypeCustom(tString)
	getTypeCustom(tBool)

}

func getTypeCustom(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("type: integer")
	case string:
		fmt.Println("type: string")
	case bool:
		fmt.Println("type: bool")
	case chan int:
		fmt.Println("type: channel int")
	case chan bool:
		fmt.Println("type: channel bool")
	case chan string:
		fmt.Println("type: channel string")
	}
}
