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
	MaxIntMayBe32 int        = 1<<63 - 1 // proved to be 64 bit on a 64 bit machine
	/*The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned
	integer type. */
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf("Test File %g \n", MaxIntMayBe32)
	fmt.Println(strconv.IntSize);
}