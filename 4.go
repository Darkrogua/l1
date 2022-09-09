package main

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C.
	Выбрать и обосновать способ завершения работы всех воркеров.
*/

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"time"
)

func main() {
	var workerCount int

	chanelSignal := make(chan os.Signal, 1)
	signal.Notify(chanelSignal, os.Interrupt)

	for setWorkerCount(&workerCount) != nil {
	}

	chanelData := make(chan int64)
	for i := 1; i <= workerCount; i++ {
		go demon(i, chanelData)
	}

	for {

		select {
		case <-chanelSignal:
			fmt.Printf("\n остановили вручную \n")
			return
		default:
			chanelData <- NewCryptoRand()
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func demon(n int, c chan int64) {
	fmt.Printf("Демон %d начал выполнение, данные: %d \n", n, <-c)
}

func setWorkerCount(c *int) error {
	fmt.Print("Введите число воркеров: ")
	_, err := fmt.Scan(c)
	return err
}

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
