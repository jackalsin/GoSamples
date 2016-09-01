package main

import "fmt"

func twoReturnTypes(x string, y string) (string, string) {
	return y, x
}

func testDeclaration() {
	var a, b, c int
	var booleanValue bool
	// wrong way to use it
	// var d = 0    d = 0, var d, int d
	d := 0
	fmt.Println(a, b, c, d, booleanValue)

}

func main() {
	a, b := twoReturnTypes("Hello", "world")
	fmt.Printf(a + b + "\n")
	testDeclaration()
}


