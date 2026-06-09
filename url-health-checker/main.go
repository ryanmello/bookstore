package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

func main() {
	args := os.Args

	var urls []string

	for i := 1; i < len(args); i++ {
		urls = append(urls, args[i])
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	
	for _, url := range urls {
		fmt.Println("Checking", url)
		// make a request to each url to check if it is active
		res, err := client.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		res.Body.Close()

		if res.StatusCode >= 200 && res.StatusCode < 300 {
			fmt.Println("Healthy:", res.Status)
		} else {
			fmt.Println("Not healthy:", res.Status)
		}
	}

	// use concurrency to check multiple urls while waiting on others
}
