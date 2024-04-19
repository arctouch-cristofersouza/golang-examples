package main

import "fmt"

//make
// -  memory allocation and initialization.
// - use for slices, maps, and channels
//func make(t Type, size ...IntegerType) Type

func main() {
	fmt.Println("-- MAKE --")
	a := make([]int, 0)
	aPtr := &a
	fmt.Println("pointer == nil :", *aPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *aPtr)

	fmt.Println("-- COMPOSITE LITERAL --")
	b := []int{}
	bPtr := &b
	fmt.Println("pointer == nil :", *bPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *bPtr)

	fmt.Println("-- NEW --")
	cPtr := new([]int)
	fmt.Println("pointer == nil :", *cPtr == nil)
	//returns a pointer to the zero value of that type
	fmt.Printf("pointer value: %p\n\n", *cPtr)

	fmt.Println("-- VAR (not initialized) --")
	var d []int
	dPtr := &d
	fmt.Println("pointer == nil :", *dPtr == nil)
	fmt.Printf("pointer value: %p\n", *dPtr)
}
