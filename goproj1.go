package main

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

func main() {
	var test = Vector{1, 2}
	rand := Math.rand(10)
	fmt.Println("Hello world!")
	fmt.Println(test)
}
