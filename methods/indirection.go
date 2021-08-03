package main

import "fmt"

type Vertex4 struct {
	X, Y float64
}

//method for Vertex4
func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//function
func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex4{3, 4}
	v.Scale(2) // interpreter v as &v; OK
	ScaleFunc(&v, 10)

	p := &Vertex4{4, 3}
	p.Scale(3) // receive p as &v; OK
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
