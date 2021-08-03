package main

//GO does not have classes, but can define methods on types
import (
	"fmt"
	"math"
)

//struct type as "class"
type Vertex struct {
	X, Y float64
}

//method for type Vertex with special argument before name Abs
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
it's usual function - not method
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
