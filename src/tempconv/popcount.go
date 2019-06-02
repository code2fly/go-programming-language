package tempconv

import "fmt"

// pc[i] is population count of i
var pc [256]byte

// init is special func that can be declared any number of times and it is executed by default in a file in order it is declared
func init() {
	fmt.Println("this should get executed first")
}

func init() {
	fmt.Println("this function should be executed by default")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func init() {
	fmt.Println("i guess this func should be executed as well not sure though")
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
