package main

import (
//  "bufio"
 "os"
 "fmt"
)

// func main() {
	
// }



// func dup_initial() {
// 	//   create a map
// 	counts := make(map[string] int)
// 	// read line from input
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// check if line already exists in map if so increament the counter
// 	for line, count := range counts {
// 		if count > 1 {
// 			fmt.Printf("%d\t%s\n", count, line)
// 		}
// 	}
// }


// this version operates in streaming mode, in which input is read and broken into lines as needed and hence it can operate on even large inputs
func duplines_init() {
	counts := make(map[string]int)
	files:= os.Args[1:]
	
	if len(files)  == 0 {
		countLines(os.Stdin, counts)
	}	else {
		for _,filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
 
}


func countLines(file *os.File , count map[string]int) {

}


