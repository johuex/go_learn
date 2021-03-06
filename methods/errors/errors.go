package main

import (
	"fmt"
	"time"
)

//A `nil` error denotes success; a `non-nil` error denotes failure.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// the `error` type is a built-in interface similar to fmt.Stringer
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
