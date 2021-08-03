package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
if we delete pointer in receiver, than we work with copy of v.
With pointer in receiver we work with original
*/
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
