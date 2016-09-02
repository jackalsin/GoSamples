package main

import "fmt"

func main() {
	useVoidPointer()

	tryVoidPtrFromStringToFloat()
}

func tryVoidPtrFromStringToFloat() {
	var i interface{} = "33"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	//hello true
	f, ok := i.(float64) // wont work.
	fmt.Println(f, ok)
	// 0 false
	f = i.(float64) // panic
	fmt.Println(f)
}

func useVoidPointer() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
