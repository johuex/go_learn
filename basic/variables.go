package main

import "fmt"

// all declared statements are automatically reset to 0, '' or false
var c, python, java bool

//but declaration can include initializing by = (not :=)
var z, l int = 1, 2

//Constants cannot be declared using the := syntax
const Pi = 3.14

func main() {
	var i int
	var golang, rust = true, "no!"
	//Inside a function, the := short assignment statement can be used
	//but outside statement begins with var and := isn't available
	k := 3
	fmt.Println(i, z, l, k, c, python, java, golang, rust, Pi)
}
