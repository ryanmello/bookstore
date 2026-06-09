package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
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

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		
		go func(url string){
			defer wg.Done()

			res, err := client.Get(url)
			if err != nil {
				ch <- fmt.Sprintf("%s is not healthy: %v", url, err)
				return
			}

			defer res.Body.Close()

			if res.StatusCode >= 200 && res.StatusCode < 300 {
				ch <- fmt.Sprintf("%s is healthy: %s", url, res.Status)
			} else {
				ch <- fmt.Sprintf("%s is not healthy: %s", url, res.Status)
			}
		}(url)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
