package main

import (
	"fmt"
)

// hash table or map
func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 4

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

//to more info https://pkg.go.dev/maps@go1.22.1
