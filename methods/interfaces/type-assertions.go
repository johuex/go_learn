package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) //example of type assertion
	fmt.Println(s)  //hello

	s, ok := i.(string)
	fmt.Println(s, ok) //hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) //0 false

	f = i.(float64) // panic, because ok isn't defined
	fmt.Println(f)
}
