package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(rand.Intn(100), rand.Intn(100)))
}
