package main

import "fmt"

var x uint8 = 1<<1 | 1<<5
var y uint8 = 0xef

func main() {

	// %b verb to print bit pattern, 08 to print as 8 bit
	fmt.Printf("Printing int as binary byte data - %08b \n", x)
	// implementation of integer sets that can be much bigger than byte
	int_set()

	fmt.Println(y)

	// decimal , octal and hexadecimal
	o := 0666
	// [1] adverb tell % to use first operand over and over again
	fmt.Printf("%o , %[1]d , %[1]x \n", o)

	// runes literal are single chars written in single quotes
	var char rune
	char = 'a'
	// %c and %q can be used while printing apart from %d
	fmt.Printf("%d , %[1]c , %[1]q ", char)
	char = 60
	fmt.Printf("%d , %[1]c , %[1]q ", char)
	char = '\n'
	// %c and %q can be used while printing apart from %d
	fmt.Printf("%d , %[1]c , %[1]q ", char)
}

func int_set() {

	var x int8 = 1<<1 | 1<<5
	var y int8 = 1<<1 | 1<<2

	fmt.Printf("%08b \n", x)	// "00100010" set of {1,5}
	fmt.Printf("%08b \n", y)    // "00000110" set of {1,2}

	fmt.Printf("%08b \n", x&y)  // intersection resutl should be {1}

	// membership test
	for i := uint8(0); i < 8 ; i++ {
		if x&(1<<i) != 0 {
			fmt.Println("membership test passed by : ", i) // passed by {1,5}
		}
	}

}
