package main

import (
	"log"
	"fmt"
)

func main() {
	log.Println("Let's see")

	// for i, j := 0, 0; i < 3; i++, j-- { ...}
	// This is illegal
	var	p *[]int = new([]int);

	fmt.Println(*p)
	v := make([]int, 10)
	fmt.Println(v)

	for i := 0; i < len(v); i++ {
		v[i] = i
	}
	sli := v[0:3]
	sli = append(sli, 11)

	fmt.Println(v)
	fmt.Println(sli)
}