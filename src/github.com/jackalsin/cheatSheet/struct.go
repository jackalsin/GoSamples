package main

import "fmt"

func main() {
	pointerExample()
	fmt.Println(Vertex{3, 5})
	mapMethod()
}

type Vertex struct {
	X int
	Y int
}

func mapMethod() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, isPresent := m["Answer"]
	fmt.Println("The value:", v, "Present?", isPresent)

	value := m["Answer"]
	fmt.Println("The value:", value)

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
