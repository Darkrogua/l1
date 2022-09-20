package main

import (
	"fmt"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	//  Отправить в канал с которым работает горутина данные, которая она примит как инструкцию к остановке

	c := make(chan int)
	q := make(chan int)

	go secondWorker(c, q)

Loop:
	for i := 100; i > 50; i-- {
		select {
		case <-q:
			close(c)
			break Loop
		default:
			c <- i
		}
	}

}

func secondWorker(c, q chan int) {
	for {
		if <-c != 70 {
			fmt.Println("Делаю работу ", <-c)
		} else {
			fmt.Println("Завершаем работу получил код 70")
			q <- 0
			break
		}
	}

}
