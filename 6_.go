package main

import (
	"fmt"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	// пристрелить мейн
	// прокинуть канал который придет с данными горутина их обработает и написать обработку если пришли эти данные то все каюк
	// контекст прокинуть (надо подумать)

	// контекст cancel
	// контекст timeout
	// контекст

	//1. Законочить выполнение программы, оставить main
	c := make(chan int)
	q := make(chan int)

	//num := 10
	//var stop bool

	//go func() {
	//	for {
	//		select {
	//		case <-c:
	//			fmt.Println("Делаю полезную работу", <-c)

	//		case <-quit:
	//			fmt.Println("Завершаем работу работу")
	//			return
	//		}
	//	}
	//}()

	//for stop == false {
	//	c <- 5
	//	timer := time.AfterFunc(time.Second*time.Duration(3), func() {
	//		stop = true
	//	})

	//	defer timer.Stop()
	//}

	//2. Отправить в канал с которым работает горутина данные, которая она примит как инструкцию к остановке

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

	// Тут надо прокинуть контекст

}
