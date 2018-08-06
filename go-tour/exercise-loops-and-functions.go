package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 0.000000000001

	for math.Abs(math.Pow(z, 2)-x) > delta {
		z = z - (math.Pow(z, 2)-x)/(2*x)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
