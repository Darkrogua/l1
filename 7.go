package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	mumbersMap := make(map[int]int)

	wg := new(sync.WaitGroup)

	var mutex sync.Mutex

	for i, n := range numbers {
		wg.Add(1)
		go setMap(mumbersMap, wg, &mutex, n, i)
	}

	wg.Wait()

	fmt.Println(mumbersMap)

}

func setMap(m map[int]int, wg *sync.WaitGroup, mutex *sync.Mutex, i, n int) {
	defer mutex.Unlock()
	defer wg.Done()
	mutex.Lock()
	m[i] = n
}
