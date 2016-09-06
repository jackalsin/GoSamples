package main

import (
	"fmt"
	"math"
)

const DELTA float64 = 1e-8;

func Sqrt(x float64) float64 {
	z := x
	prevZ := -1.0
	for ; math.Abs(z - prevZ) > DELTA; {
		prevZ = z
		z = z - (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

