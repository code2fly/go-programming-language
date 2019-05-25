package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

// this is a sample UNIX echo command impl
func main()  {

	var res  string;
	if len(os.Args) > 1 {
		for i := 1 ; i < len(os.Args) ; i++ {
			res += os.Args[i] + " "
		}
		echo_forConditional(os.Args[1:]) 
		// echo_forRange(os.Args[1:])
		echo_StringsJoin(os.Args)
	}

	// fmt.Println(res)

}


// echo variation with only conditional for loop
func echo_forConditional(args []string)  {
	startTime := time.Now()
	var index int
	var res string
	for index < len(args) {
		res += res + " "
		index++
	}
	fmt.Println("time taken in for conditional ===> ", time.Since(startTime).Nanoseconds())
	fmt.Println("Printing from for conditional ===> ", res)
}


// echo variation with range for loop ** see how this works as results are weird
func echo_forRange(args []string) {
	var res string
	fmt.Println("printing from range function")
	for index, value := range args {
		fmt.Println("index : ", index, " value", value)
		res += res + value + " "
	}
	fmt.Println("Final Printing from for range ===> ", res)
}


// better performant echo rather than appending string of slice
func echo_StringsJoin(args []string) {
	startTime := time.Now()
	var res string
	res = strings.Join(args, " ")

	fmt.Println("time taken in for strings.join ===> ", time.Since(startTime).Nanoseconds())
	fmt.Println("final Print from join function ===> ", res)
}