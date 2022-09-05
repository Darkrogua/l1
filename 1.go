package main

import (
	"fmt"
	"time"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func main() {
	arnold := Human{"Арни", 75}

	arnold.sayName()
	fmt.Println(arnold.getYearBorn())

	action := Action{arnold}

	action2 := Action{
		Human: Human{
			name: "Boris",
			age:  23,
		},
	}

	fmt.Println(action.name)

	fmt.Println(action2.name)
}

func (h *Human) sayName() {
	fmt.Println(h.name)
}

func (h *Human) getYearBorn() int {
	year, _, _ := time.Now().Date()
	return year - 75
}
