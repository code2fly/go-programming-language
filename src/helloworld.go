package main

import (
	"fmt"
	"log"
	"tempconv"
)

var cwd string

func main() {
	// go natively handles Unicode, so it can process text in all the world's languages
	fmt.Println("hello हर कोई")
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.PopCount(65))
	fmt.Println(5.0 / 2)
	var x int8 = -128
	fmt.Println(uint8(x))
	fmt.Println(2 & 3 , 2 | 3 , 2 ^ 3 , 2 &^ 3)
}

func init() {
	var err error
	if err != nil {
		log.Fatalf("os.Getwd failed with error : %v", err)
	} else {
		log.Println("current working directory :  ", cwd)
	}
}
