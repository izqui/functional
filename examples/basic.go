package main

import (
	"fmt"
	f "github.com/izqui/functional"
	"github.com/izqui/helpers"
)

func main() {

	words := []string{"doing", "go", "functional", "programming"}
	numbers := []int{1, 2, 3}

	fmt.Println("MAP")

	fmt.Println("Square", numbers, f.DoMap(func(x int) int { return x * x }, numbers))
	fmt.Println("Length", words, f.DoMap(func(x string) int { return len(x) }, words))
	fmt.Println("Hash", words, f.DoMap(hash, words))

	/*
		fmt.Println("REDUCE")

		fmt.Println("Sum", numbers, f.Reduce(func(a int, x int) int { return a + x }, numbers, 0))
		fmt.Println("Total length", words, f.Reduce(func(a interface{}, x interface{}) interface{} { return a.(int) + len(x.(string)) }, words, 0))
	*/

	fmt.Println("FILTER")
	fmt.Println("Long words (>5)", words, f.Filter(func(x string) bool { return (len(x) > 5) }, words))

}

func hash(x string) string {

	return helpers.SHA1([]byte(x))
}
