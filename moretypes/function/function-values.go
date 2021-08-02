package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	// there we give function-as-parametr two arguments
	return fn(3, 4)
}

func main() {
	//Functions are values too.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// there we hand over functions as arguments in function compute
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
