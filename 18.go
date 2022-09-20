package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

var counter int

func main() {
	fmt.Printf("Счетчик: %d", counter)
	done := make(chan struct{})

	var mutex = sync.Mutex{}

	go worker(1, &mutex, done)
	go worker(2, &mutex, done)
	go worker(3, &mutex, done)

	time.Sleep(4 * time.Second)

	close(done)

	fmt.Printf("Счётчик: %d", counter)
}

func worker(workerId int, m *sync.Mutex, done chan struct{}) {
	for {
		select {
		case _, ok := <-done:
			if !ok {
				return
			}
		default:
		}

		m.Lock()
		counter++
		fmt.Printf("Worker %d, counter %d\n", workerId, counter)

		m.Unlock()

		time.Sleep(200 * time.Millisecond)
	}
}
