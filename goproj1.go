package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Vector : creates a two-dimensional vector
type Vector struct {
	X int
	Y int
}

func midpoint(p1, p2 Vector) (mid Vector) {
	mid = Vector{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2} 
	return
}

func generateRandomMatrix(width, height int) [][]int {
	matrix := make([][]int, width)
	for i:= range matrix {
		matrix[i] = make([]int, height)
	}

	rand.Seed( time.Now().UTC().UnixNano() )
	
	for i := 0; i < width; i++ {
		for j:= 0; j < width; j++ {
			matrix[i][j] = rand.Int() % 10
		}
	}

	return matrix
}

func main() {
	var test = Vector{1, 2}
	point1 := Vector{6,12}
	point2 := Vector{0,0}

	fmt.Println("Hello World!")
	fmt.Println(midpoint(point1, point2))
	fmt.Println(test)
	fmt.Println(generateRandomMatrix(3,3))
}
