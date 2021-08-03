package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //here we closed a channel because no more values will be sent
	//Only the sender should close a channel, never the receiver. But don't usually need to close
	//them. Closing is only necessary when the receiver must be told there are no more values coming,
	//such as to terminate a range loop.
}

func main() {
	c := make(chan int, 11)
	go fibonacci(cap(c), c) //cap(c) - capacity (not length) of channel c
	for i := range c {      //we close c for ranging in for-loop
		fmt.Println(i)
	}
}
