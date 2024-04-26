package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

// Recover is a way to threat panic problems
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
