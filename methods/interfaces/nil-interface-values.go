package main

import "fmt"

type I4 interface {
	M()
}

func main() {
	var i I4
	describe3(i)
	i.M() //there is no type inside the interface tuple to indicate which concrete method to call.
}

func describe3(i I4) {
	fmt.Printf("(%v, %T)\n", i, i)
}
