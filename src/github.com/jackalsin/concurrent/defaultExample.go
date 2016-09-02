package main

import "fmt"

func main()  {
	f("Direct")
	// It's like push a task to a deque, like nodejs express, infinitely looping from the beginning to the end to check
	// if this is available to execute
	go f("Go routine")

	go func(msg string) {
		fmt.Println(msg + ": Calliing going")
	}("Anonymous")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func f(str string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":" ,i)
	}
}