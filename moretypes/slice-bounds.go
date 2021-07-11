package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[0:] //The default is zero for the low bound
	// and the length of the slice for the last high bound (for this example = 2)
	//so s[0:] = s[0:2]
	fmt.Println(s)
}
