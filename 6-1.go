package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	//1. Законочить выполнение программы, оставить main
	c := make(chan int)
	q := make(chan int)

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("Делаю полезную работу", <-c)
			case <-q:
				fmt.Println("Завершаем работу работу")
				return
			}
		}
	}()

	var stop bool

	for stop == false {
		c <- 5
		timer := time.AfterFunc(time.Second*time.Duration(2), func() {
			stop = true
		})

		defer timer.Stop()
	}

}
