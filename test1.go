package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Vector : creates a two-dimensional vector of float64 values
type Vector struct {
	X, Y float64
}

func midpoint(p1, p2 Vector) (mid Vector) {
	mid = Vector{(p1.X + p2.X) / 2, (p1.Y + p2.Y) / 2}
	return
}

func generateMatrix(width, height int) [][]int {
	matrix := make([][]int, width)
	for i:= range matrix {
		matrix[i] = make([]int, height)
	}

	return matrix
}

func generateProblemMatrix() (matrix [3][3]int) {

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Int() % 10
		}
	}

	return
}

func fibonacciRecursive(maxRange , n1, n2 int) {
	defer fmt.Println(n1)

	maxRange--

	if maxRange == 1 {
		return
	}

	fibonacciRecursive(maxRange, n2, n2 + n1)

}

func fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		a, b = a + b, a
		return b
	}
}

type integer int

func (i integer) Add(n integer) integer {
	return i + n
}

func main() {
	var test = Vector{1, 2}
	point1 := Vector{6,12}
	point2 := Vector{0,0}
	f := fibonacci()

	fmt.Println("Hello World!")
	fmt.Println(midpoint(point1, point2))
	fmt.Println(test)
	fmt.Println(generateMatrix(3,3))
	fmt.Println(generateProblemMatrix())
	fmt.Println(math.Sqrt(2), math.Sqrt2)

	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
	//fibonacciRecursive(100, 0, 1)

	fmt.Println(integer(10).Add(10))
	fmt.Println(pow(10, 0), pow(1, 100), pow(7,7)) //Error at pow(7, 7)
}
