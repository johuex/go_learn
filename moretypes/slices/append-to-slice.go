package main

import "fmt"

func main() {
	var s []int
	printSlice_2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice_2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice_2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice_2(s)
}

func printSlice_2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
