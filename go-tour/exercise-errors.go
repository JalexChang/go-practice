package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	delta := 0.000000000001
	for math.Abs(math.Pow(z, 2)-x) > delta {
		z = z - (math.Pow(z, 2)-x)/(2*x)
	}

	return z, nil
}

func printSqrt(x float64) {
	if result, err := Sqrt(x); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
	return
}

func main() {
	printSqrt(2)
	printSqrt(-2)
}
