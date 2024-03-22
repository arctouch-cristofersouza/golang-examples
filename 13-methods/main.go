package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) speak() {
	fmt.Println("Hello from " + p.Name)
}

func main() {
	flavio := Person{Age: 39, Name: "Flavio"}
	flavio.speak()
}
