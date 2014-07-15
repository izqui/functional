package main

import (
	"fmt"
)

type AnonFunc func(x interface{}) interface{}
type AnonSlice []interface{}

func main() {

	numbers := AnonSlice{1, 2, 3}
	fmt.Println("Square", numbers, doMap(func(x interface{}) interface{} { return x.(int) * x.(int) }, numbers))

	words := AnonSlice{"hola", "que", "tal", "estas"}
	fmt.Println("Length", words, doMap(func(x interface{}) interface{} { return len(x.(string)) }, words))
}

func doMap(f AnonFunc, vs AnonSlice) (r AnonSlice) {

	r = make(AnonSlice, len(vs))
	for i, v := range vs {
		r[i] = f(v)
	}
	return r
}
