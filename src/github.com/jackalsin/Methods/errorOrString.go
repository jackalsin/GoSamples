package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) String() string {
	return "calling String()"
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run1() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func run2() *MyError {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run1(); err != nil {
		fmt.Println(err)
	}

	if err := run2(); err != nil {
		fmt.Println(err)
	}
}
