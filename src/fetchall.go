package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// get the url
	urls := os.Args[1:]
	// create a channel
	channel := make(chan string)
	// execute fetch goroutine
	for _, url := range urls {
		go fetch_concurrent(url, channel) // start go routine
	}

	for range urls {
		fmt.Println(<-channel) // recieve from channel
	}
	fmt.Printf("total time for executing the main Goroutine ===> %.2fs", time.Since(start).Seconds())

}

func fetch_concurrent(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stdout, "error while fetching data %v\n", err)
		os.Exit(1)
	}

	// data , error := ioutil.ReadAll(resp.Body)
	inputSize, error := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if error != nil {
		fmt.Fprintf(os.Stdout, "error while reading data from response body %v\n", err)
	}

	ch <- fmt.Sprintf("%.2fs time for data size from : %v is ====> %v\n", time.Since(start).Seconds(), url, inputSize)

}
