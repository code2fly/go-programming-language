package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {

	for _, url := range os.Args[1:] {

		data := fetch_data(url)
		fmt.Fprintf(os.Stdout, "data retrieved in response ======> %s\n ", string(data))
	}

}


func fetch_data(url string) string {

	resp, err := http.Get(url)

	if err != nil  {
		fmt.Fprintf(os.Stdout, "error while fetching data %v\n", err)
		os.Exit(1)
	}
	
	data , error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if error != nil {
		fmt.Fprintf(os.Stdout, "error while reading data from response body %v\n", err)
	}
	
	return string(data)
}