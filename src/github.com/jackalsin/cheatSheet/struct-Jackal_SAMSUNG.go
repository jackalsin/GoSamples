package main

import "fmt"

func main() {
	pointerExample()
	fmt.Println(Vertex{3, 5})

	// access with a pointer
	v := Vertex{1, 2}
	p := &v
	// this is right
	p.X = 1e9
	fmt.Println(v)
	(*p).X = 1e11
	fmt.Println(v)

	// array

	// slice
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	sliceDefault()
	sliceAndCapacity()

}

// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func sliceAndCapacity()  {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
type Vertex struct {
	X int
	Y int
}

func pointerExample() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p) // print the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
