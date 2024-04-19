package main

import "fmt"

//  Structs are typed collections of fields. They’re useful for grouping data together to form records.

// struct
type Person struct {
	Name string
	Age  int
}

// embedd struct
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	flavio := Person{"Flavio", 39}

	mario := Person{Age: 40, Name: "Mario"}

	fmt.Println(mario)
	fmt.Println(flavio)

	//embedd
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
