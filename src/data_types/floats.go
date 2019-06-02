package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	fmt.Println(math.NaN)

	for i:=0 ; i < 8 ; i ++ {
		// %g verb can be used to print the most compact form
		fmt.Printf("i : %d , e^i : %8.3e  \n", i , math.Exp(float64(i)))
	}


}


func checkNan(num float64) (float64, error){
	if math.IsNaN(num) {
		return float64(0) , fmt.Errorf("this is not a number")
	}
	return num, nil
}


func surface() {
	
}