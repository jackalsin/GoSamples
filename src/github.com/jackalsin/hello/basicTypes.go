package main

import (
	"fmt"
	"math/cmplx"
	"strconv"
	"math"
	"time"
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

	// requires explicitly converse
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(-fl)
	fmt.Println(x, y, z)

	// const: you cannot declare
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// for loop
	//  the braces { } are always required.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statement are optional
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	// while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)


	t := time.Now()
	fmt.Println("Time Now " + t.String())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	deferUsage()
	// print 987654321
	deferIsStacked()
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// it ranges from index and value.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	skipIndex()

}
func skipIndex() {

	// skip index or value
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
func deferIsStacked() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func deferUsage()  {
	// defer will be only executed when returns.
	defer fmt.Println("world")
	fmt.Println("hello")
}
func pow(x, n, lim float64) float64 {
	// can start with a short statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}