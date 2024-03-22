package main

import "fmt"

const globalConstant = 3.14

func main() {
	var thisIs string
	thisIs = "Isso é"

	oneType := "uma string"

	fmt.Println(thisIs, oneType)
	//fmt.Print(globalConstant)
	//fmt.Printf("%s %s e isso um float %.2f ", thisIs, oneType, globalConstant)
}
