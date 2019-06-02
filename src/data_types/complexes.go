package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// create a complex number
	var x complex128 = complex(1, 2)
	var y = complex(2, -3)
	fmt.Println(x)
	fmt.Println(x * y)

	// we can also write
	z := 1 + 3i
	fmt.Println(z)

	// complex library function like square root
	fmt.Println(cmplx.Sqrt(-1))
}



func mandelbrot() {
	
}