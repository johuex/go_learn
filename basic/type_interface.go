package main

import "fmt"

func main() {
	//When the right hand side of the declaration is typed,
	//the new variable is of that same type
	v := 42 // change me to int, float and etc
	fmt.Printf("v is of type %T\n", v)
}
