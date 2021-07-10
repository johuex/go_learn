package main

import "fmt"

func main() {
	sum := 0
	sum_con := 1
	for i := 0; i < 10; i++ { //variables declared there (for example, i)
		// are visible only in the scope of the for statement.
		sum += i
	}
	/*
	Go has only one looping construct, the for loop.
	For using "while loop" in Go u can drop semicolon:
	*/
	for ; sum_con < 1000; { // or for sum_con < 1000
		sum_con += sum_con
	}
	// Forever "for loop" looks like:
	//for {
	// }

	fmt.Println(sum)
	fmt.Println(sum_con)
}
