package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise")

	fmt.Println(ex002(2000, 3200))
}

func ex002(low, high int) string {
	res := []string{}
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 > 0 {
			value := strconv.Itoa(i)
			res = append(res, value)
		}
	}
	return strings.Join(res, ",")
}

// Create a unit test
func fib(n int) []string {
	res := []string{}
	p := 0
	c := 0
	for j := 1; c < n; j += p {
		value := strconv.Itoa(p)
		res = append(res, value)
		p = j - p
		c++
	}
	return res
}
