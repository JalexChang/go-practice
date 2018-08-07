package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	delta := 0.0000000001
	z := complex128(1)
	for cmplx.Abs(cmplx.Pow(z, 3)-x) > delta {
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(Cbrt(4))
	fmt.Println(Cbrt(8))
}
