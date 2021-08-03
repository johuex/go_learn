package main

import "fmt"

/*
Channels are a typed conduit through which you can send and receive values with the channel operator `<-`
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//sums the numbers in a slice, distributing the work between two goroutines
	//on each goroutine has half of the array
	c := make(chan int)     //create channel
	go sum(s[:len(s)/2], c) //create first goroutine
	go sum(s[len(s)/2:], c) //create second goroutine
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
