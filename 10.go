package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	temperaturs := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	var sortTemp = make(map[int][]float32)

	for _, t := range temperaturs {
		sortTemp[getKeyTemp(t)] = append(sortTemp[getKeyTemp(t)], t)
	}

	fmt.Println(sortTemp)

}

func getKeyTemp(t float32) int {
	return int(t) / 10 * 10
}
