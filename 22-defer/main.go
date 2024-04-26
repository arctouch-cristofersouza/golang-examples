package main

import (
	"fmt"
)

func whatIstheValue(val int) {
	fmt.Println("Value is", val)
}

func main() {

	defer whatIstheValue(1)

	fmt.Println("end of main")
}
