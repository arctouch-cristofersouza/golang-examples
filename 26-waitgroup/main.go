package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(d time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // don't forget to sinalize that you finish
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("Work is done")
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(2)
	go doWork(time.Second*2, &wg)
	go doWork(time.Second*4, &wg)

	wg.Wait()

	fmt.Println("Work took", time.Since(start))
}
