package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtdWorkers := 3

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}
}
