package main

import "fmt"

func main() {

	var name = "test"
	//Verify the length of a string
	fmt.Println(len(name))

	//Search by position //// ASCII
	fmt.Println("\nPegar char em uma posicao")
	fmt.Println(name[0])
	fmt.Println(string(name[2])) //"test"

	//Get chat at position
	fmt.Println("\nPegar partes da string")
	fmt.Println(name[0:2]) //"test"
	fmt.Println(name[2:])  //"test"

	//copy
	var newstring = name[:]
	fmt.Printf("Original %s e copia %s\n", name, newstring)

	//Always a Copy
	var otherString = "another one"
	oneMorePlease := otherString
	fmt.Printf("Original %s e copia %s\n", otherString, oneMorePlease)
	otherString = "new string"
	fmt.Printf("Original %s e copia %s\n", otherString, oneMorePlease)

}

//to more functions https://pkg.go.dev/strings
