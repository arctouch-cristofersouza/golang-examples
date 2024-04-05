package main

import "fmt"

func main() {
	a := 1    // var a int
	b := 3.14 // var b float
	c := "hi" // var c string
	d := true // var d bool
	fmt.Printf("Temos os seguintes tipos, um int %d, um float %.2f, uma string %s e um boolean %t \n", a, b, c, d)

	// e is a SLICE, not a array
	e := []int{1, 2, 3, 4}
	fmt.Println(e, len(e), e[0], e[1:3], e[1:], e[:2])

	//F is a array
	var f [2]string
	f[0] = "Hello"
	f[1] = "World"
	fmt.Println(f[0], f[1], len(f))

	// OBS:
	// the difference between a SLICE and ARRAY is that a slice append new items on the end and array have a FIXED leng
	//

	//G is a map
	g := make(map[string]int)
	g["one"] = 1
	g["two"] = 2
	fmt.Println(g, len(g), g["one"], g["three"])
}
