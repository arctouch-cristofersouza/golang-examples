package main

import (
	"fmt"
	"strings"
)

func plus(a int, b int) int {
	return a + b
}

// Forma diferente de declarar os parametros
func plusPlus(a, b, c int) int {
	return a + b + c
}

func splitWords(word, delimeter string) (string, string) {
	words := strings.Split(word, delimeter)
	return words[0], words[1]
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//multiples returns
	word1, word2 := splitWords("one - two", "-")
	fmt.Println(word1, word2)

	_, word3 := splitWords("one - three", "-")
	fmt.Println(word1, word2, word3)

}
