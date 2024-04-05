package main

import "fmt"

func main() {
	//For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//While
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//with Break
	k := 0
	for {
		fmt.Println(k)

		if k < 10 {
			break
		}

		i++
	}

}
