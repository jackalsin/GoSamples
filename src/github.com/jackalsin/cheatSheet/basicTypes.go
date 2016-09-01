package main

import (
	"fmt"
	"math/cmplx"
	"strconv"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	//MaxIntMayBe32 int        = 1 << 64 - 1
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	//fmt.Printf("Test File %g \n", MaxIntMayBe32)
	fmt.Println(strconv.IntSize);
}