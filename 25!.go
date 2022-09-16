package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleepTimeAfter(t time.Duration) {
	_ = <-time.After(t)
}

func main() {
	fmt.Println("Запуск сна!")

	sleepTimeAfter(2 * time.Second)
	fmt.Println("Проснулись")

}
