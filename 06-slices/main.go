package main

import "fmt"

func main() {
	var mySlice []string

	otherSlice := []string{"First", "Second", "Third"}

	fmt.Printf("Tamanho do meu slice é %d\n", len(mySlice))
	fmt.Printf("Tamanho do meu slice é %d\n", len(otherSlice))

	otherSlice = append(otherSlice, "Fourth", "Fifth")
	fmt.Printf("Adiconado mais dados no slice, agora tem tamanho de %d\n", len(otherSlice))

	makeASlice := make([]string, 3) //Fazendo um slice com 3 strings vazias

	fmt.Printf("Tamanho do meu slice é %d\n", len(makeASlice))

}

//to more details and func https://pkg.go.dev/slices@go1.22.1
