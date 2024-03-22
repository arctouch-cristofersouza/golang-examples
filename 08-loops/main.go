package main

import "fmt"

func main() {
	//default one
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//with Break
	i = 0
	for {
		fmt.Println(i)

		if i < 10 {
			break
		}

		i++
	}

	//with range
	numbers := []int{1, 2, 3}

	for i, num := range numbers {
		fmt.Printf("%d: %d\n", i, num)
	}

}
