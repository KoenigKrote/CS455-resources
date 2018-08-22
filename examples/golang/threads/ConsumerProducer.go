package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Go emphasizes the use of channels, which combined with goroutines makes consumer/producer relationships very simple.
	Channels are built-in threadsafe queues, more or less.
*/

// The <- indicates a one-way relationship to a channel
// <-chan means "Only read"
// chan<- means "Only write"
// chan means "Read/Write"
func Consumer(inChan <-chan int) {
	for x := range inChan {
		fmt.Println("Received", x, "from channel")
	}
}

func Producer(wg *sync.WaitGroup, outChan chan<- int, start, end int) {
	for i := start; i < end; i++ {
		outChan <- i
		time.Sleep(100)
	}
	wg.Done()
}

func main() {
	// Make a channel with capacity 100
	// Channels will block on read when empty, and block on write when full
	// Channels will also accept interface types, pointers, structs, or even other channels
	channel := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Consumer(channel)
		go Producer(&wg, channel, i*5, (i*5)+5)
	}
	// Synchronized WaitGroup, wait for all producers to finish
	wg.Wait()

	// Close channel, breaks the infinite read loop for Consumers
	close(channel)

	fmt.Println("All Producers finished, channel closed, all consumers finished")
}
