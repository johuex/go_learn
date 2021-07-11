package main

import "fmt"

func main() {
	fmt.Println("counting")
	// defer pushes function into stack, so that functions will be call from
	//stack before end of function where described (in our case - in main)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
