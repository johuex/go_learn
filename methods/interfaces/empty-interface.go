package main

import "fmt"

/*
Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Print takes any number of arguments of type interface{}.
*/
func main() {
	var i interface{} // empty interface may hold values of any type.
	describe4(i)

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
