package main

import "fmt"

func main() {

	age := 17
	if age < 18 {
		fmt.Println("Voce não tem permissão para beber e dirigir no brasil")
	}

	age = 19
	if age < 18 {
		fmt.Println("Voce não tem permissão para beber e dirigir no brasil")
	} else {
		fmt.Println("Voce é adulto legalmente no brasil")
	}

	age = 11
	if age < 12 {
		fmt.Println("Voce é uma criança")
	} else if age < 18 {
		fmt.Println("Voce não tem permissão para beber e dirigir no brasil")
	} else {
		fmt.Println("Voce é adulto legalmente no brasil")
	}

	//switch

	switch age {
	case 0:
		fmt.Println("Zero years old")
	case 1:
		fmt.Println("One year old")
	case 2:
		fmt.Println("Two years old")
	case 3:
		fmt.Println("Three years old")
	case 4:
		fmt.Println("Four years old")
	default:
		fmt.Println("%i years old", age)
	}

}
