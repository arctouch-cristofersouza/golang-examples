package main

import (
	"fmt"
	"time"
)

func doWork(d time.Duration) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("Work is done")
}

func main() {
	start := time.Now()

	doWork(time.Second * 2)
	doWork(time.Second * 4)

	//go doWork(time.Second * 2)
	//go doWork(time.Second * 4)

	fmt.Println("Work took", time.Since(start))
}
