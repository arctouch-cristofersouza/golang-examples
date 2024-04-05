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

func sumElipsis(nums ...int) int {

	res := 0

	for _, n := range nums {
		res += n
	}

	return res
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//multiples returns
	word1, word2 := splitWords("one - two", "-")
	fmt.Println(word1, word2)

	//Blank Identifier
	_, word3 := splitWords("one - three", "-")
	fmt.Println(word1, word2, word3)

	//anonymous
	sum := func(a, b, c int) int {
		return a + b + c
	}(3, 5, 7)

	fmt.Println("5+3+7 =", sum)

	//elipsis
	s1 := sumElipsis(1, 2, 3)
	s2 := sumElipsis(1, 2, 3, 4)
	s3 := sumElipsis(1, 2, 3, 4, 5)

	fmt.Println(s1, s2, s3)

}
