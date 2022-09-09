package main

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	var N int = 5
	chanelData := make(chan int64)

	go demonRead(chanelData)

	for N > 0 {
		number := NewCryptoRand()
		fmt.Printf("Отправляем в канал число: %d \n", number)
		chanelData <- number
		time.Sleep(300 * time.Millisecond)

		timer := time.AfterFunc(time.Second*time.Duration(N), func() {
			fmt.Println("ВРЕМЯ")
			N = 0
		})
		defer timer.Stop()
	}
}

func demonRead(c chan int64) {
	for {
		fmt.Printf("Демон читает число из канала: %d \n", <-c)
	}
}

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
