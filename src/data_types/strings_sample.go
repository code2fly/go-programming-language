package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {

	// basic_string()
	// literals_string()
	unicode_samples()

}

// basic string decleration
func basic_string() {
	s := "Hello , world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[3])
	// 32 is the difference between unicode of capital and small case
	fmt.Printf("%c , %c \n", s[0]+32, s[3])
	// below results in a panic
	// fmt.Println(s[len(s)])

}

func literals_string() {

	// this works since GO's source file is also interpreted as utf-8
	// this is double quoted string literal
	s1 := "hello हर कोई!"
	fmt.Println(s1)
	s1 += "\t i am fine btw!! \n"
	fmt.Println(s1)

	// now lets see raw string literal using ``
	s2 := `this is a new thing and this can \t even 
extend over a few lines and it does not matter frankly \n wont even matter now `
	fmt.Println(s2)
}

func unicode_samples() {
	s := "Hello, 世界 , हर कोई!"
	// normal and clumsy way
	// for len(s) > 0 {
	// 	r, size := utf8.DecodeRuneInString(s)
	// 	fmt.Printf("unicode char is : %[1]c, code point value is: %[1]v , size is %[2]d \n", r, size)
	// 	s = s[size:]
	// }

	// better way using range , when applied over string, performs UTF-8 decoding implicitly
	for index, value := range s {
		fmt.Printf("unicode char is : %[1]q, code point value is: %[1]v , currnet index is : %[2]d \n", value, index)
	}

	// []rune conversion on a string
	r := []rune(s)
	fmt.Println(r)

	// int to string conversion
	fmt.Println(string(65), string(1234567))

}
