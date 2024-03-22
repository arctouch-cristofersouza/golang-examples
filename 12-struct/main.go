package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	flavio := Person{"Flavio", 39}

	mario := Person{Age: 40, Name: "Mario"}

	fmt.Println(mario)
	fmt.Println(flavio)

}
