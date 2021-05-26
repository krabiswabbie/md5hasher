package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Worker(in <-chan string, out chan<- string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for v := range in {
		// Make request
		resp, err := client.Get(v)
		if err != nil {
			out <- fmt.Sprintf("%v returned error: %v", v, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			out <- fmt.Sprintf("%v returned %v status code", v, resp.StatusCode)
			continue
		}

		// Request successful
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			out <- fmt.Sprintf("Failed to read from %v: %v", v, err)
			continue
		}
		resp.Body.Close()

		md := md5.Sum(b)
		out <- fmt.Sprintf("%v %v", v, hex.EncodeToString(md[:]))
	}
}

func checkParameters() (int, []string) {
	p := flag.Int("parallel", 10, "Parallel requests")
	flag.Parse()
	return *p, flag.Args()
}

func main() {
	workers, addrs := checkParameters()

	in := make(chan string, len(addrs))
	out := make(chan string, len(addrs))

	// Launch workers
	for i := 0; i < workers; i++ {
		go Worker(in, out)
	}

	// Transfer jobs
	for _, v := range addrs {
		in <- v
	}

	// Wait all workers to complete requests
	for i := 0; i < len(addrs); i++ {
		log.Println(<-out)
	}

	close(in)
}
