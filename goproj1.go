package main

import (
	"fmt"
	"math/rand"
)

type Vector struct {
	X int
	Y int
}

func main() {
	var test = Vector{1, 2}
	rand := rand.Intn(100)
	fmt.Println("Hello world!")
	fmt.Println(test, rand)
}
