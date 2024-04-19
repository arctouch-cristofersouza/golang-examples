package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// / Pointers on Methods
type Employee struct {
	name string
	age  int
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

type rectangle struct {
	length int
	width  int
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	fmt.Println("========================")

	//pointers on method
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Println("Employee name before change:", e.name)
	e.changeName("Michael Andrew")
	fmt.Println("Employee name after change:", e.name)

	fmt.Println("Employee age before change:", e.age)

	person := &e
	person.changeAge(51)
	fmt.Println("Employee age after change:", e.age)

	///// error, Access a Nil pointer will crash you program
	// var p *int
	// fmt.Println(*p)

}

// - https://dev.to/diwakarkashyap/pointers-in-golang-go-1910
