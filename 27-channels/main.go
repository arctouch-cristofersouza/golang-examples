package main

import (
	"fmt"
	"time"
)

// Channels and go routines (Golang Concurrency)

// Channels is use to manage concurrency on golang

func main() {
	ch := make(chan int) // Creating an unbuffered channel
	ch2 := make(chan int)

	// goroutine is a function that is capable of running concurrently with other functions
	go func() {
		fmt.Println("Doing work on channel 1")
		time.Sleep(time.Second) // Simulating some work
		ch <- 5                 // Sending a value to the channel
	}()

	go func() {
		fmt.Println("Doing work on channel 2")
		time.Sleep(time.Second * 3)
		ch2 <- 10
	}()

	x := <-ch      // Receiving the value from the channel
	fmt.Println(x) // Output: 5

	y := <-ch2
	fmt.Println(y)

}

func closeChannel() {
	ch := make(chan int) // Creating an unbuffered channel

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Sending values to the channel
		}
		close(ch) // Closing the channel after all values are sent
	}()

	for x := range ch {
		fmt.Println(x) // Output: 1, 2, 3, 4, 5
	}
}

// benefits of channels is their ability to synchronize goroutines.
// By using channels, you can coordinate the execution of
// multiple goroutines, ensuring that they complete their work.

// Go article to read about it: https://medium.com/goturkiye/concurrency-in-go-channels-and-waitgroups-25dd43064d1
