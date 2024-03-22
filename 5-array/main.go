package main

import "fmt"

func main() {
	var myArray [3]string
	myArray[2] = "Another"

	preFilledArray := [3]string{"First", "Second", "Third"}
	var goCountArray = [...]string{"First", "Second", "Third"}

	//Print first position
	fmt.Println(myArray[0])
	fmt.Println(preFilledArray[0])
	fmt.Println(goCountArray[0])

	//What is the length?
	fmt.Println(len(myArray))
	fmt.Println(len(preFilledArray))
	fmt.Println(len(goCountArray))

}
